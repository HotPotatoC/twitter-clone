BEGIN;
ALTER TABLE IF EXISTS favorites
    DROP CONSTRAINT IF EXISTS favorites_id_tweet_id_user_unique;
ALTER TABLE IF EXISTS favorites
    DROP CONSTRAINT IF EXISTS favorites_id_tweet_foreign;
ALTER TABLE IF EXISTS favorites
    DROP CONSTRAINT IF EXISTS favorites_id_user_foreign;
DROP TABLE IF EXISTS favorites;
COMMIT;

