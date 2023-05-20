-- name: CreateUser :one
INSERT INTO users (name, mail, status) VALUES ($1, $2, $3)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: ListUserByTask :one
SELECT u.* FROM users u INNER JOIN user_task_relation r ON r.user_id = u.id WHERE r.task_id = $1;

-- name: UpdateUserAsLoggedIn :one
UPDATE users SET status = 2 WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;
