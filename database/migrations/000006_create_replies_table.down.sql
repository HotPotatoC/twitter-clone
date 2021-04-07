BEGIN;
ALTER TABLE IF EXISTS replies
    DROP CONSTRAINT IF EXISTS replies_id_tweet_unique;
ALTER TABLE IF EXISTS replies
    DROP CONSTRAINT IF EXISTS replies_id_tweet_foreign;
DROP TABLE IF EXISTS replies;
COMMIT;

