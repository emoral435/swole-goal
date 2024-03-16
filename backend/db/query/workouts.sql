-- CreateWorkout: returns a new workout, provided uid, title, body, and last time modified/used
--
-- returns: the new workout row
-- name: CreateWorkout :one
INSERT INTO "workouts" (
    user_id, title, body, last_time
) VALUES (
    (SELECT "users".id FROM "users" WHERE "users".id = $1), $2, $3, $4
) RETURNING *;

-- GetUserWorkouts: returns a users workouts, provided their uid
--
-- returns: the user's corresponding workout rows
-- name: GetUserWorkouts :many
SELECT * FROM "workouts"
WHERE user_id = $1;

-- GetNumWorkouts: returns the number of workouts a user has, provided their uid
--
-- returns: the user's corresponding workout rows
-- name: GetNumWorkouts :one
SELECT COUNT(*) FROM "workouts"
WHERE user_id = $1;

-- GetWorkout: returns an existing workout, given workout id
--
-- returns: the corresponding workout row
-- name: GetWorkout :one
SELECT * FROM "workouts"
WHERE id = $1 AND user_id = $2 LIMIT 1;

-- UpdateWorkoutTitle: updates workouts title given its id
--
-- returns: the workout's new corresponding row
-- name: UpdateWorkoutTitle :one
UPDATE "workouts"
SET title = $3
WHERE id = $1 AND user_id = $2
RETURNING *;

-- UpdateBody: updates workout's body text given its workouts id
--
-- returns: the workouts new row
-- name: UpdateWorkoutBody :one
UPDATE "workouts"
SET body = $3
WHERE id = $1 AND user_id = $2
RETURNING *;

-- UpdateLastWorkout: updates workout's last workout time given its id
--
-- returns: the workout's new corresponding row
-- name: UpdateWorkoutLast :one
UPDATE "workouts"
SET last_time = $3
WHERE id = $1 AND user_id = $2
RETURNING *;

-- DeleteSingleWorkout: deletes a single user's workout
--
-- returns: nothing! see https://docs.sqlc.dev/en/stable/reference/query-annotations.html for exec
-- name: DeleteSingleWorkout :exec
DELETE FROM "workouts"
WHERE id = $1 AND user_id = $2;

-- DeleteAllWorkouts: deletes All user's workouts
--
-- returns: nothing! see https://docs.sqlc.dev/en/stable/reference/query-annotations.html for exec
-- name: DeleteAllWorkouts :exec
DELETE FROM "workouts"
WHERE user_id = $1;
