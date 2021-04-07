BEGIN;
ALTER TABLE IF EXISTS likes
    DROP CONSTRAINT IF EXISTS likes_id_tweet_id_user_unique;
ALTER TABLE IF EXISTS likes
    DROP CONSTRAINT IF EXISTS likes_id_tweet_foreign;
ALTER TABLE IF EXISTS likes
    DROP CONSTRAINT IF EXISTS likes_id_user_foreign;
DROP TABLE IF EXISTS likes;
COMMIT;

