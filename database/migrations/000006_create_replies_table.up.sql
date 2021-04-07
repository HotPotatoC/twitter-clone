BEGIN;
CREATE TABLE IF NOT EXISTS "replies" (
    "id_reply" bigint,
    "id_tweet" integer NOT NULL
);
ALTER TABLE "replies"
    ADD CONSTRAINT "replies_id_reply_unique" UNIQUE ("id_reply");
ALTER TABLE "replies"
    ADD CONSTRAINT "replies_id_tweet_foreign" FOREIGN KEY ("id_tweet") REFERENCES "tweets" ("id") MATCH FULL ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE "replies"
    ADD PRIMARY KEY ("id_reply");
CREATE INDEX "replies_id_tweet_index" ON "replies" ("id_tweet");
COMMIT;

