-- CreateSet: returns a new set, providing:
-- a exercise id, and reps and weights
--
-- returns: the new set row
-- name: CreateSet :one
INSERT INTO "set" (
    exercise_id, reps, weight
) VALUES (
    (SELECT "workouts".id FROM "workouts" WHERE "workouts".id = $1), $2, $3
) RETURNING *;

-- GetExerciseSets: returns all exercises sets
--
-- returns: the exercises corresponding sets
-- name: GetExerciseSets :many
SELECT * FROM "set"
WHERE exercise_id = $1;

-- GetSet: returns an existing Set, given Set id
--
-- returns: the corresponding Set row
-- name: GetSet :one
SELECT * FROM "set"
WHERE id = $1 LIMIT 1;

-- UpdateSetRep: updates sets reps given its id
--
-- returns: the set's new corresponding row
-- name: UpdateSetRep :one
UPDATE "set"
SET reps = $2
WHERE id = $1
RETURNING *;

-- UpdateSetWeight: updates set's weight given its id
--
-- returns: the sets's new corresponding row
-- name: UpdateSetWeight :one
UPDATE "set"
SET weight = $2
WHERE id = $1
RETURNING *;

-- DeleteSingleSet: deletes a single exercises's set
--
-- returns: nothing! see https://docs.sqlc.dev/en/stable/reference/query-annotations.html for exec
-- name: DeleteSingleSet :exec
DELETE FROM "set"
WHERE id = $1;

-- DeleteAllSets: deletes All exercises sets!
--
-- returns: nothing! see https://docs.sqlc.dev/en/stable/reference/query-annotations.html for exec
-- name: DeleteAllSets :exec
DELETE FROM "set"
WHERE exercise_id = $1;
