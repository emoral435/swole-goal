-- CreateUser: returns a new user, provided their email, password, and username
--
-- returns: the new user row
-- name: CreateUser :one
INSERT INTO "users" (
    email, password, username
) VALUES (
    $1, $2, $3
) RETURNING *;

-- GetUser: returns a new user, provided their uid
--
-- returns: the user's corresponding row
-- name: GetUser :one
SELECT * FROM "users"
WHERE id = $1 LIMIT 1;

-- NumUsers: returns the number of users
--
-- returns: the number of users
-- name: NumUsers :one
SELECT COUNT(*) FROM "users";

-- GetUserEmail: returns an existing user, given their unique email
--
-- returns: the user's corresponding row
-- name: GetUserEmail :one
SELECT * FROM "users"
WHERE email = $1 LIMIT 1;

-- ListUsers: returns all users in the database
--
-- returns: all users
-- name: ListUsers :many
SELECT * FROM "users"
ORDER BY id
LIMIT
    1
    OFFSET 2;

-- UpdatePassword: updates user's password given their uid
--
-- returns: the user's new corresponding row
-- name: UpdatePassword :one
UPDATE "users"
SET password = $2
WHERE id = $1
RETURNING *;

-- UpdatePasswordEmail: updates user's password given their email
--
-- returns: the user's new corresponding row
-- name: UpdatePasswordEmail :one
UPDATE "users"
SET password = $2
WHERE email = $1
RETURNING *;

-- UpdateUsername: updates user's username given their uid
--
-- returns: the user's new corresponding row
-- name: UpdateUsername :one
UPDATE "users"
SET username = $2
WHERE id = $1
RETURNING *;

-- UpdateUsernameEmail: updates user's username given their email
--
-- returns: the user's new corresponding row
-- name: UpdateUsernameEmail :one
UPDATE "users"
SET username = $2
WHERE email = $1
RETURNING *;

-- UpdateBirthday: updates user's birthday given their uid
--
-- returns: the user's new corresponding row
-- name: UpdateBirthday :one
UPDATE "users"
SET birthday = $2
WHERE id = $1
RETURNING *;


-- DeleteUser: deletes a user given their uid
--
-- returns: nothing! see https://docs.sqlc.dev/en/stable/reference/query-annotations.html for exec
-- name: DeleteUser :exec
DELETE FROM "users"
WHERE id = $1;
