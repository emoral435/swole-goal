CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "username" varchar UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "birthday" timestamptz
);

CREATE TABLE "workout" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "user_id" bigserial NOT NULL,
  "title" varchar NOT NULL,
  "body" text NOT NULL,
  "last" timestamp NOT NULL,
  "exe1" bigserial,
  "exe2" bigserial,
  "exe3" bigserial,
  "exe4" bigserial,
  "exe5" bigserial,
  "exe6" bigserial,
  "exe7" bigserial
);

CREATE TABLE "exercise" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "workout_id" bigserial NOT NULL,
  "type" varchar NOT NULL,
  "title" varchar NOT NULL,
  "desc" text,
  "set1" bigint,
  "weight1" bigint,
  "set2" bigint,
  "weight2" bigint,
  "set3" bigint,
  "weight3" bigint,
  "set4" bigint,
  "weight4" bigint,
  "last_volume" bigint NOT NULL DEFAULT 0
);

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "workout" ("user_id");

CREATE INDEX ON "workout" ("exe1");

CREATE INDEX ON "workout" ("exe2");

CREATE INDEX ON "workout" ("exe3");

CREATE INDEX ON "workout" ("exe4");

CREATE INDEX ON "workout" ("exe5");

CREATE INDEX ON "workout" ("exe6");

CREATE INDEX ON "workout" ("exe7");

COMMENT ON COLUMN "users"."email" IS 'email to sign in - also to send reminders';

COMMENT ON COLUMN "workout"."body" IS 'Description of workout';

COMMENT ON COLUMN "workout"."last" IS 'Timestamp of the last time completed';

COMMENT ON COLUMN "exercise"."type" IS 'The body section this exercise hits - chest, back, etc.';

COMMENT ON COLUMN "exercise"."title" IS 'What is the exercise called?';

COMMENT ON COLUMN "exercise"."desc" IS 'description of the exercies - good for reminders';

COMMENT ON COLUMN "exercise"."last_volume" IS 'tracks what the overall volume was the last time this exercise was performed';

ALTER TABLE "workout" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "workout" ADD FOREIGN KEY ("exe1") REFERENCES "exercise" ("id");

ALTER TABLE "workout" ADD FOREIGN KEY ("exe2") REFERENCES "exercise" ("id");

ALTER TABLE "workout" ADD FOREIGN KEY ("exe3") REFERENCES "exercise" ("id");

ALTER TABLE "workout" ADD FOREIGN KEY ("exe4") REFERENCES "exercise" ("id");

ALTER TABLE "workout" ADD FOREIGN KEY ("exe5") REFERENCES "exercise" ("id");

ALTER TABLE "workout" ADD FOREIGN KEY ("exe6") REFERENCES "exercise" ("id");

ALTER TABLE "workout" ADD FOREIGN KEY ("exe7") REFERENCES "exercise" ("id");

ALTER TABLE "exercise" ADD FOREIGN KEY ("workout_id") REFERENCES "workout" ("id");
