package ctrl

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/matheus-alvs01dev/go-boilerplate/internal/app/api/schema"
	"github.com/matheus-alvs01dev/go-boilerplate/internal/domain/entity"
	"github.com/pkg/errors"
)

type UserService interface {
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
	GetByID(ctx context.Context, uuid uuid.UUID) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, uuid uuid.UUID) error
}

type UserController struct {
	service UserService
}

func NewUserController(svc UserService) *UserController {
	return &UserController{
		service: svc,
	}
}

func (uc *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var req schema.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		http.Error(w, errors.Wrap(err, "validation").Error(), http.StatusBadRequest)
		return
	}

	user := entity.NewUser(req.Name, req.Email, req.Wallet)
	createdUser, err := uc.service.Create(r.Context(), user)
	if err != nil {
		http.Error(w, errors.Wrap(err, "svc").Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(schema.NewCreateUserResponse(createdUser))
}

func (uc *UserController) Update(w http.ResponseWriter, r *http.Request) {
	var req schema.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		http.Error(w, errors.Wrap(err, "validation").Error(), http.StatusBadRequest)
		return
	}

	userID := uuid.MustParse(req.ID)
	user := entity.NewUser(req.Name, req.Email, req.Wallet)
	user.ID = userID

	if err := uc.service.Update(r.Context(), user); err != nil {
		http.Error(w, errors.Wrap(err, "svc").Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (uc *UserController) GetByID(w http.ResponseWriter, r *http.Request) {
	userIDStr := chi.URLParam(r, "id")

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		http.Error(w, schema.NewValidationError("id is invalid").Error(), http.StatusBadRequest)
		return
	}

	user, err := uc.service.GetByID(r.Context(), userID)
	if err != nil {
		http.Error(w, errors.Wrap(err, "svc").Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (uc *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	userIDStr := chi.URLParam(r, "id")

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		http.Error(w, schema.NewValidationError("id is invalid").Error(), http.StatusBadRequest)
		return
	}

	if err := uc.service.Delete(r.Context(), userID); err != nil {
		http.Error(w, errors.Wrap(err, "svc").Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
