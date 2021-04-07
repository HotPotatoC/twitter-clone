BEGIN;
ALTER TABLE IF EXISTS follows
    DROP CONSTRAINT IF EXISTS follows_follower_id_foreign;
ALTER TABLE IF EXISTS follows
    DROP CONSTRAINT IF EXISTS follows_followed_id_foreign;
ALTER TABLE IF EXISTS tweets
    DROP CONSTRAINT IF EXISTS tweets_user_id_foreign;
DROP TABLE IF EXISTS users;
COMMIT;

