BEGIN;
ALTER TABLE tweets
    ADD COLUMN content_tsv tsvector;
UPDATE
    tweets
SET
    content_tsv = setweight(to_tsvector(content), 'A');
CREATE INDEX content_index ON tweets USING GIN (to_tsvector('english', content));
CREATE OR REPLACE FUNCTION tweet_tsvector_trigger ()
    RETURNS TRIGGER
    AS $$
BEGIN
    NEW.content_tsv := setweight(to_tsvector('english', NEW.content), 'A');
    RETURN NEW;
END
$$
LANGUAGE plpgsql;
CREATE TRIGGER trigger_tweets_tsvector
    BEFORE INSERT OR UPDATE ON tweets
    FOR EACH ROW
    EXECUTE PROCEDURE tweet_tsvector_trigger ();
COMMIT;

