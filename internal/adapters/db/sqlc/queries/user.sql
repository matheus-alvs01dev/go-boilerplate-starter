-- name: CreateUser :exec
INSERT INTO users (name, email, wallet)
VALUES ($1, $2, $3);

-- name: GetUserByID :one
SELECT
    id,
    name,
    email,
    wallet,
    created_at,
    updated_at
FROM users
WHERE id = $1
  AND deleted_at IS NULL;

-- name: UpdateUser :exec
UPDATE users
SET
    name       = $1,
    email      = $2,
    wallet     = $3,
    updated_at = NOW()
WHERE id = $4
  AND deleted_at IS NULL;

-- name: DeleteUser :exec
UPDATE users
SET deleted_at = NOW()
WHERE id = $1
  AND deleted_at IS NULL;
