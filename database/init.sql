/* source: https://drawsql.app/templates/twitter */
CREATE TABLE "follows"(
    "id" BIGINT GENERATED ALWAYS AS IDENTITY,
    "follower_id" INTEGER NULL,
    "followed_id" INTEGER NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
ALTER TABLE "follows"
ADD CONSTRAINT "follows_follower_id_followed_id_unique" UNIQUE("follower_id", "followed_id");
ALTER TABLE "follows"
ADD PRIMARY KEY("id");
CREATE INDEX "follows_follower_id_index" ON "follows"("follower_id");
CREATE INDEX "follows_followed_id_index" ON "follows"("followed_id");
CREATE TABLE "tweets"(
    "id" BIGINT GENERATED ALWAYS AS IDENTITY,
    "content" VARCHAR(255) NULL,
    "user_id" INTEGER NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
CREATE INDEX "tweets_user_id_created_at_index" ON "tweets"("user_id", "created_at");
ALTER TABLE "tweets"
ADD PRIMARY KEY("id");
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