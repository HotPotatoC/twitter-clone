BEGIN;
CREATE TABLE "users"(
    "id" BIGINT GENERATED ALWAYS AS IDENTITY,
    "name" VARCHAR(255) NULL,
    "email" VARCHAR(255) NULL,
    "password" VARCHAR(255) NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
ALTER TABLE "users"
ADD PRIMARY KEY("id");
ALTER TABLE "users"
ADD CONSTRAINT "users_email_unique" UNIQUE("email");
ALTER TABLE "follows"
ADD CONSTRAINT "follows_follower_id_foreign" FOREIGN KEY("follower_id") REFERENCES "users"("id");
ALTER TABLE "follows"
ADD CONSTRAINT "follows_followed_id_foreign" FOREIGN KEY("followed_id") REFERENCES "users"("id");
ALTER TABLE "tweets"
ADD CONSTRAINT "tweets_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("id") ON DELETE CASCADE;
COMMIT;