BEGIN;
ALTER TABLE IF EXISTS tweets
    DROP CONSTRAINT IF EXISTS "tweets_user_id_created_at_index";
ALTER TABLE IF EXISTS users
    DROP CONSTRAINT IF EXISTS "tweets_user_id_foreign";
ALTER TABLE IF EXISTS tweets RENAME COLUMN user_id TO id_user;
CREATE INDEX "tweets_id_user_created_at_index" ON "tweets" ("id_user", "created_at");
ALTER TABLE "tweets"
    ADD CONSTRAINT "tweets_id_user_foreign" FOREIGN KEY ("id_user") REFERENCES "users" ("id") ON DELETE CASCADE;
COMMIT;

