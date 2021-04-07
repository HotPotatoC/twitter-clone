BEGIN;
CREATE TABLE IF NOT EXISTS "users" (
    "id" bigint GENERATED ALWAYS AS IDENTITY,
    "name" varchar(255) NULL,
    "email" varchar(255) NULL,
    "password" varchar(255) NULL,
    "created_at" timestamp(0) without time zone NOT NULL
);
ALTER TABLE "users"
    ADD PRIMARY KEY ("id");
ALTER TABLE "users"
    ADD CONSTRAINT "users_email_unique" UNIQUE ("email");
ALTER TABLE "follows"
    ADD CONSTRAINT "follows_follower_id_foreign" FOREIGN KEY ("follower_id") REFERENCES "users" ("id");
ALTER TABLE "follows"
    ADD CONSTRAINT "follows_followed_id_foreign" FOREIGN KEY ("followed_id") REFERENCES "users" ("id");
ALTER TABLE "tweets"
    ADD CONSTRAINT "tweets_user_id_foreign" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;
COMMIT;

