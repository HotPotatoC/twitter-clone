BEGIN;
CREATE TABLE IF NOT EXISTS "tweets" (
    "id" bigint GENERATED ALWAYS AS IDENTITY,
    "content" varchar(255) NULL,
    "user_id" integer NULL,
    "created_at" timestamp(0) without time zone NOT NULL
);
CREATE INDEX "tweets_user_id_created_at_index" ON "tweets" ("user_id", "created_at");
ALTER TABLE "tweets"
    ADD PRIMARY KEY ("id");
COMMIT;

