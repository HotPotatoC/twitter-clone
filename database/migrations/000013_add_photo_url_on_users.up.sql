BEGIN;
ALTER TABLE IF EXISTS "users"
    ADD COLUMN photo_url TEXT NULL;
UPDATE
    users
SET
    photo_url = 'https://twitterclone-bucket.s3.amazonaws.com/default-avatar.png';
COMMIT;

