BEGIN;
CREATE TABLE IF NOT EXISTS "retweets" (
    "id_user" bigint NOT NULL,
    "id_tweet" bigint NOT NULL,
    "created_at" timestamp(0) without time zone NOT NULL
);
ALTER TABLE "retweets"
    ADD PRIMARY KEY ("id_user", "id_tweet");
ALTER TABLE "retweets"
    ADD CONSTRAINT "retweets_id_tweet_foreign" FOREIGN KEY ("id_tweet") REFERENCES "tweets" ("id") MATCH FULL ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE "retweets"
    ADD CONSTRAINT "retweets_id_user_foreign" FOREIGN KEY ("id_user") REFERENCES "users" ("id") MATCH FULL ON UPDATE CASCADE ON DELETE CASCADE;
COMMIT;

