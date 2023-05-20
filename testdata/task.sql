-- name: CreateTask :one
INSERT INTO tasks (name, expiration_date, is_done) VALUES ($1, $2, $3)
RETURNING *;

-- name: ListTask :many
SELECT * FROM tasks ORDER BY id DESC;