-- CreateExercise: returns a new exercise, providing:
-- a workout id, type of exercise (chest, back, etc.), title, desc of exercise, and the last overall volume
--
-- returns: the new exercise row
-- name: CreateExercise :one
INSERT INTO "exercises" (
    workout_id, type, title, description, last_volume
) VALUES (
    (SELECT "workouts".id FROM "workouts" WHERE "workouts".id = $1), $2, $3, $4, $5
) RETURNING *;

-- GetWorkoutsExercise: returns all workouts exercises
--
-- returns: the workouts corresponding exercise rows
-- name: GetWorkoutsExercise :many
SELECT * FROM "exercises"
WHERE workout_id = $1;

-- GetExercise: returns an existing exercise, given exercise id
--
-- returns: the corresponding exercise row
-- name: GetExercise :one
SELECT * FROM "exercises"
WHERE id = $1 LIMIT 1;

-- UpdateExerciseType: updates exercises type given its id
--
-- returns: the exercise's new corresponding row
-- name: UpdateExerciseType :one
UPDATE "exercises"
SET type = $2
WHERE id = $1
RETURNING *;

-- UpdateExerciseTitle: updates exercises title given its id
--
-- returns: the exercise's new corresponding row
-- name: UpdateExerciseTitle :one
UPDATE "exercises"
SET title = $2
WHERE id = $1
RETURNING *;

-- UpdateExerciseDescription: updates exercise's description text given its exercises id
--
-- returns: the exercises new row
-- name: UpdateExerciseDescription :one
UPDATE "exercises"
SET description = $2
WHERE id = $1
RETURNING *;

-- UpdateLastExercise: updates exercise's last exercise time given its id
--
-- returns: the exercise's new corresponding row
-- name: UpdateExerciseLast :one
UPDATE "exercises"
SET last_volume = $2
WHERE id = $1
RETURNING *;

-- DeleteSingleExercise: deletes a single workouts's exercise
--
-- returns: nothing! see https://docs.sqlc.dev/en/stable/reference/query-annotations.html for exec
-- name: DeleteSingleExercise :exec
DELETE FROM "exercises"
WHERE id = $1;

-- DeleteAllExercises: deletes All workout's exercises
--
-- returns: nothing! see https://docs.sqlc.dev/en/stable/reference/query-annotations.html for exec
-- name: DeleteAllExercises :exec
DELETE FROM "exercises"
WHERE workout_id = $1;
