BEGIN;
CREATE TABLE "tweets"(
    "id" BIGINT GENERATED ALWAYS AS IDENTITY,
    "content" VARCHAR(255) NULL,
    "user_id" INTEGER NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
CREATE INDEX "tweets_user_id_created_at_index" ON "tweets"("user_id", "created_at");
ALTER TABLE "tweets"
ADD PRIMARY KEY("id");
COMMIT;