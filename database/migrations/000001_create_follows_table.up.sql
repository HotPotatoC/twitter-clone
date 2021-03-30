BEGIN;
CREATE TABLE IF NOT EXISTS "follows"(
    "id" BIGINT GENERATED ALWAYS AS IDENTITY,
    "follower_id" INTEGER NULL,
    "followed_id" INTEGER NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
ALTER TABLE "follows"
ADD CONSTRAINT "follows_follower_id_followed_id_unique" UNIQUE("follower_id", "followed_id");
ALTER TABLE "follows"
ADD PRIMARY KEY("id");
CREATE INDEX "follows_follower_id_index" ON "follows"("follower_id");
CREATE INDEX "follows_followed_id_index" ON "follows"("followed_id");
COMMIT;