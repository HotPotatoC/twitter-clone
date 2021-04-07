BEGIN;
CREATE TABLE IF NOT EXISTS "likes" (
    "id" bigint GENERATED ALWAYS AS IDENTITY,
    "id_tweet" integer NULL,
    "id_user" integer NULL,
    "created_at" timestamp(0) without time zone NOT NULL
);
ALTER TABLE "likes"
    ADD CONSTRAINT "likes_id_tweet_id_user_unique" UNIQUE ("id_tweet", "id_user");
ALTER TABLE "likes"
    ADD CONSTRAINT "likes_id_tweet_foreign" FOREIGN KEY ("id_tweet") REFERENCES "tweets" ("id") ON DELETE CASCADE;
ALTER TABLE "likes"
    ADD CONSTRAINT "likes_id_user_foreign" FOREIGN KEY ("id_user") REFERENCES "users" ("id") ON DELETE CASCADE;
ALTER TABLE "likes"
    ADD PRIMARY KEY ("id");
CREATE INDEX "likes_id_tweet_index" ON "likes" ("id_tweet");
CREATE INDEX "likes_id_user_index" ON "likes" ("id_user");
COMMIT;

