BEGIN;
ALTER TABLE IF EXISTS likes
    DROP CONSTRAINT IF EXISTS likes_id_tweet_id_user_unique;
ALTER TABLE IF EXISTS likes
    DROP CONSTRAINT IF EXISTS likes_id_tweet_foreign;
ALTER TABLE IF EXISTS likes
    DROP CONSTRAINT IF EXISTS likes_id_user_foreign;
ALTER TABLE IF EXISTS "likes" RENAME TO "favorites";
ALTER TABLE IF EXISTS "likes_id_seq" RENAME TO "favorites_id_seq";
ALTER TABLE "favorites"
    ADD CONSTRAINT "favorites_id_tweet_id_user_unique" UNIQUE ("id_tweet", "id_user");
ALTER TABLE "favorites"
    ADD CONSTRAINT "favorites_id_tweet_foreign" FOREIGN KEY ("id_tweet") REFERENCES "tweets" ("id") ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE "favorites"
    ADD CONSTRAINT "favorites_id_user_foreign" FOREIGN KEY ("id_user") REFERENCES "users" ("id") ON UPDATE CASCADE ON DELETE CASCADE;
ALTER INDEX IF EXISTS "likes_pkey" RENAME TO "favorites_pkey";
ALTER INDEX IF EXISTS "likes_id_tweet_index" RENAME TO "favorites_id_tweet_index";
ALTER INDEX IF EXISTS "likes_id_user_index" RENAME TO "favorites_id_user_index";
COMMIT;

