BEGIN;
ALTER TABLE IF EXISTS "follows"
    DROP CONSTRAINT IF EXISTS "follows_follower_id_foreign";
ALTER TABLE IF EXISTS "follows"
    DROP CONSTRAINT IF EXISTS "follows_followed_id_foreign";
ALTER TABLE "follows"
    ADD CONSTRAINT "follows_follower_id_foreign" FOREIGN KEY ("follower_id") REFERENCES "users" ("id") MATCH FULL ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE "follows"
    ADD CONSTRAINT "follows_followed_id_foreign" FOREIGN KEY ("followed_id") REFERENCES "users" ("id") MATCH FULL ON UPDATE CASCADE ON DELETE CASCADE;
COMMIT;

