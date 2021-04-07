BEGIN;
DROP INDEX IF EXISTS content_index;
DROP TRIGGER IF EXISTS trigger_tweets_tsvector ON tweets;
DROP FUNCTION IF EXISTS tweet_tsvector_trigger;
ALTER TABLE IF EXISTS tweets
    DROP COLUMN content_tsv;
COMMIT;

