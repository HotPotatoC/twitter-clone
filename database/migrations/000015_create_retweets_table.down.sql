BEGIN;
ALTER TABLE IF EXISTS retweets
    DROP CONSTRAINT IF EXISTS retweets_id_user_foreign;
ALTER TABLE IF EXISTS retweets
    DROP CONSTRAINT IF EXISTS retweets_id_tweet_foreign;
DROP TABLE IF EXISTS retweets;
COMMIT;

