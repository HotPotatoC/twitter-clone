BEGIN;
ALTER TABLE IF EXISTS "users"
    ADD COLUMN handle VARCHAR(255) NULL;
ALTER TABLE "users"
    ADD CONSTRAINT "users_handle_unique" UNIQUE ("handle");
UPDATE
    users
SET
    handle = replace(name, ' ', '');
COMMIT;

