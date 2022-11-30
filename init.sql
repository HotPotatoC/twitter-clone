BEGIN;

CREATE TABLE IF NOT EXISTS users (
    "id" bigint PRIMARY KEY,
    "name" varchar NOT NULL,
    "screen_name" varchar NOT NULL,
    "email" varchar NOT NULL,
    "bio" varchar NOT NULL,
    "location" varchar NOT NULL,
    "website" varchar NOT NULL,
    "profile_image_url" text NOT NULL,
    "profile_banner_url" text NOT NULL,
    "followers_count" int NOT NULL,
    "followings_count" int NOT NULL,
    "created_at" timestamp(0) without time zone NOT NULL,
    "updated_at" timestamp(0) without time zone NOT NULL
);

CREATE TABLE IF NOT EXISTS followers (
    "followee_id" bigint NOT NULL REFERENCES users ("id") ON DELETE CASCADE,
    "follower_id" bigint NOT NULL REFERENCES users ("id") ON DELETE CASCADE,
    "created_at" timestamp(0) without time zone NOT NULL,
    PRIMARY KEY ("followee_id", "follower_id")
);

CREATE TABLE IF NOT EXISTS tweets (
    "id" bigint PRIMARY KEY,
    "user_id" bigint NOT NULL REFERENCES users ("id") ON DELETE CASCADE,
    "content" varchar(280) CHECK (char_length("content") <= 280),
    "favorites_count" int,
    "replies_count" int,
    "created_at" timestamp(0) without time zone NOT NULL
);

CREATE INDEX IF NOT EXISTS tweets_created_at_idx ON tweets ("created_at");

CREATE TABLE IF NOT EXISTS tweet_entities (
    "tweet_id" bigint REFERENCES tweets ON DELETE CASCADE,
    "media_links" text[] CHECK (array_length("media_links", 1) <= 4),
    "created_at" timestamp(0) without time zone NOT NULL
);

CREATE TABLE IF NOT EXISTS replies (
    "tweet_id" bigint NOT NULL REFERENCES tweets ("id") ON DELETE CASCADE,
    "reply_id" bigint NOT NULL REFERENCES tweets ("id") ON DELETE CASCADE,
    PRIMARY KEY ("tweet_id", "reply_id")
);

CREATE TABLE IF NOT EXISTS retweets (
    "tweet_id" bigint NOT NULL REFERENCES tweets ("id") ON DELETE CASCADE,
    "retweet_id" bigint NOT NULL REFERENCES tweets ("id") ON DELETE CASCADE,
    PRIMARY KEY ("tweet_id", "retweet_id")
);

CREATE TABLE IF NOT EXISTS favorites (
    "user_id" bigint NOT NULL REFERENCES users ("id") ON DELETE CASCADE,
    "tweet_id" bigint NOT NULL REFERENCES tweets ("id") ON DELETE CASCADE,
    PRIMARY KEY ("tweet_id", "user_id")
);

CREATE TABLE IF NOT EXISTS feeds (
    "id" bigint PRIMARY KEY,
    "user_id" bigint NOT NULL REFERENCES users ("id") ON DELETE CASCADE,
    "created_at" timestamp(0) without time zone NOT NULL
);

CREATE TABLE IF NOT EXISTS feed_tweets (
    "tweet_id" bigint NOT NULL REFERENCES tweets ("id") ON DELETE CASCADE,
    "feed_id" bigint NOT NULL REFERENCES feeds ("id") ON DELETE CASCADE,
    PRIMARY KEY ("tweet_id", "feed_id")
);

COMMIT;

