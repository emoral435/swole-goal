CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "username" varchar UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "birthday" timestamptz
);

CREATE TABLE "workouts" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "user_id" bigserial NOT NULL,
  "title" varchar NOT NULL,
  "body" text NOT NULL,
  "last_total_volume" timestamp NOT NULL
);

CREATE TABLE "exercises" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "workout_id" bigserial NOT NULL,
  "type" varchar NOT NULL,
  "title" varchar NOT NULL,
  "desc" text,
  "last_volume" bigint NOT NULL DEFAULT 0
);

CREATE TABLE "set" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "exercise_id" bigserial NOT NULL,
  "reps" bigint,
  "weight" bigint,
  "last_volume" bigint NOT NULL DEFAULT 0
);

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "workouts" ("user_id");

COMMENT ON COLUMN "users"."email" IS 'email to sign in - also to send reminders';

COMMENT ON COLUMN "workouts"."body" IS 'Description of workout';

COMMENT ON COLUMN "workouts"."last_total_volume" IS 'Timestamp of the last time completed';

COMMENT ON COLUMN "exercises"."type" IS 'The body section this exercise hits - chest, back, etc.';

COMMENT ON COLUMN "exercises"."title" IS 'What is the exercise called?';

COMMENT ON COLUMN "exercises"."desc" IS 'description of the exercies - good for reminders';

COMMENT ON COLUMN "exercises"."last_volume" IS 'tracks what the overall volume was the last time this exercise was performed';

COMMENT ON COLUMN "set"."last_volume" IS 'tracks what the overall volume was the last time this exercise was performed';

ALTER TABLE "workouts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "exercises" ADD FOREIGN KEY ("workout_id") REFERENCES "workouts" ("id");

ALTER TABLE "set" ADD FOREIGN KEY ("exercise_id") REFERENCES "exercises" ("id");
