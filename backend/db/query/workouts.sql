-- CreateWorkout: returns a new workout, provided their email, password, and username
--
-- returns: the new workout row
-- name: CreateWorkout :one
INSERT INTO "workouts" (
    user_id, title, body, last
) VALUES (
    (SELECT id FROM "users" WHERE id = $1), $2, $3, $4
) RETURNING *;

-- GetUserWorkouts: returns a users workouts, provided their uid
--
-- returns: the user's corresponding workout rows
-- name: GetUserWorkouts :many
SELECT * FROM "workouts"
WHERE uid = $1;

-- GetWorkout: returns an existing workout, given workout id
--
-- returns: the corresponding workout row
-- name: GetUserEmail :one
SELECT * FROM "workouts"
WHERE id = $1 LIMIT 1;

-- UpdateTitle: updates workouts title given its id
--
-- returns: the workout's new corresponding row
-- name: UpdateTitle :one
UPDATE "workouts"
SET title = $2
WHERE id = $1
RETURNING *;

-- UpdateBody: updates workout's body text given its workouts id
--
-- returns: the workouts
-- name: UpdateBody :one
UPDATE "workouts"
SET body = $2
WHERE id = $1
RETURNING *;

-- UpdateLastWorkout: updates workout's last workout time given its id
--
-- returns: the workout's new corresponding row
-- name: UpdateLastWorkout :one
UPDATE "workouts"
SET last = $2
WHERE id = $1
RETURNING *;

-- DeleteSingleWorkout: deletes a single user's workout
--
-- returns: nothing! see https://docs.sqlc.dev/en/stable/reference/query-annotations.html for exec
-- name: DeleteSingleWorkout :exec
DELETE FROM "workouts"
WHERE id = $1;

-- DeleteAllWorkouts: deletes All user's workouts
--
-- returns: nothing! see https://docs.sqlc.dev/en/stable/reference/query-annotations.html for exec
-- name: DeleteAllWorkouts :exec
DELETE FROM "workouts"
WHERE user_id = $1;
