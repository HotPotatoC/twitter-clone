BEGIN;
ALTER TABLE IF EXISTS "users"
    DROP COLUMN bio,
    DROP COLUMN location,
    DROP COLUMN website,
    DROP COLUMN birth_date;
COMMIT;

