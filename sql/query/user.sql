-- name: CreateUser :one
INSERT INTO users (
    user_name,
    hashed_password,
    full_name,
    email
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE user_name = $1 LIMIT 1;

-- -- name: UpdateUser :one
-- UPDATE users
--     SET 
--         hashed_password = $1,
--         email = $2
-- WHERE user_name = $3
-- RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE user_name = $1;