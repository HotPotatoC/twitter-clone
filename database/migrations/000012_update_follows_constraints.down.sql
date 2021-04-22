BEGIN;
ALTER TABLE "follows"
    DROP CONSTRAINT IF EXISTS "follows_follower_id_foreign";
ALTER TABLE "follows"
    DROP CONSTRAINT IF EXISTS "follows_followed_id_foreign";
COMMIT;

