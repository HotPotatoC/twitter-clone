BEGIN;
CREATE TABLE IF NOT EXISTS "replies"(
    "id" BIGINT GENERATED ALWAYS AS IDENTITY,
    "id_tweet" INTEGER NULL
);
ALTER TABLE "replies"
ADD CONSTRAINT "replies_id_tweet_unique" UNIQUE("id_tweet");
ALTER TABLE "replies"
ADD CONSTRAINT "replies_id_tweet_foreign" FOREIGN KEY ("id_tweet") REFERENCES "tweets"("id") MATCH FULL ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE "replies"
ADD PRIMARY KEY("id");
CREATE INDEX "replies_id_tweet_index" ON "replies"("id_tweet");
COMMIT;