BEGIN;
ALTER TABLE IF EXISTS tweets
    DROP CONSTRAINT IF EXISTS tweets_id_user_created_at_index;
ALTER TABLE IF EXISTS tweets
    DROP CONSTRAINT IF EXISTS tweets_id_user_foreign;
DROP TABLE IF EXISTS tweets;
COMMIT;

