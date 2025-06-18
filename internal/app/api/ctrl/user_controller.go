package ctrl

import (
	"context"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/matheus-alvs01dev/go-boilerplate/internal/app/api/schema"
	"github.com/matheus-alvs01dev/go-boilerplate/internal/domain/entity"
	"github.com/pkg/errors"
	"net/http"
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

func (uc *UserController) Create(c echo.Context) error {
	var req schema.CreateUserRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := req.Validate(); err != nil {
		return errors.Wrap(err, "validation")
	}

	user := entity.NewUser(req.Name, req.Email, req.Wallet)
	createdUser, err := uc.service.Create(c.Request().Context(), user)
	if err != nil {
		return errors.Wrap(err, "svc")
	}

	return c.JSON(http.StatusCreated, schema.NewCreateUserResponse(createdUser))
}

func (uc *UserController) Update(c echo.Context) error {
	var req schema.UpdateUserRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := req.Validate(); err != nil {
		return errors.Wrap(err, "validation")
	}

	userID := uuid.MustParse(req.ID)
	user := entity.NewUser(req.Name, req.Email, req.Wallet)
	user.ID = userID

	if err := uc.service.Update(c.Request().Context(), user); err != nil {
		return errors.Wrap(err, "svc")
	}

	return c.NoContent(http.StatusNoContent)
}

func (uc *UserController) GetByID(c echo.Context) error {
	userIDStr := c.Param("id")

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return schema.NewValidationError("id is invalid")
	}

	user, err := uc.service.GetByID(c.Request().Context(), userID)
	if err != nil {
		return errors.Wrap(err, "svc")
	}

	return c.JSON(http.StatusOK, user)
}

func (uc *UserController) Delete(c echo.Context) error {
	userIDStr := c.Param("id")

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return schema.NewValidationError("id is invalid")
	}

	if err := uc.service.Delete(c.Request().Context(), userID); err != nil {
		return errors.Wrap(err, "svc")
	}

	return c.NoContent(http.StatusNoContent)
}
