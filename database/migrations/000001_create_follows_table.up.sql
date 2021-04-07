BEGIN;
CREATE TABLE IF NOT EXISTS "follows" (
    "id" bigint GENERATED ALWAYS AS IDENTITY,
    "follower_id" integer NULL,
    "followed_id" integer NULL,
    "created_at" timestamp(0) without time zone NOT NULL
);
ALTER TABLE "follows"
    ADD CONSTRAINT "follows_follower_id_followed_id_unique" UNIQUE ("follower_id", "followed_id");
ALTER TABLE "follows"
    ADD PRIMARY KEY ("id");
CREATE INDEX "follows_follower_id_index" ON "follows" ("follower_id");
CREATE INDEX "follows_followed_id_index" ON "follows" ("followed_id");
COMMIT;

