-- ############################################
-- #            Seed 'Users' table            #
-- ############################################
-- All users have the same password (123123)
INSERT INTO users(name, email, password, created_at)
VALUES(
        'twitter_user1',
        'twitter_user1@email.com',
        '$2y$12$/HayGBMevM7EkbRho3xlC.nu/8.OeHGNi9iqGneNIDBrqnkD9fa1.',
        now()
    ) ON CONFLICT DO NOTHING
RETURNING id AS twitter_user1_id;
INSERT INTO users(name, email, password, created_at)
VALUES(
        'twitter_user2',
        'twitter_user2@email.com',
        '$2y$12$/HayGBMevM7EkbRho3xlC.nu/8.OeHGNi9iqGneNIDBrqnkD9fa1.',
        now()
    ) ON CONFLICT DO NOTHING
RETURNING id AS twitter_user2_id;
INSERT INTO users(name, email, password, created_at)
VALUES(
        'twitter_user3',
        'twitter_user3@email.com',
        '$2y$12$/HayGBMevM7EkbRho3xlC.nu/8.OeHGNi9iqGneNIDBrqnkD9fa1.',
        now()
    ) ON CONFLICT DO NOTHING
RETURNING id AS twitter_user3_id;
--
INSERT INTO users(name, email, password, created_at)
VALUES(
        'twitter_user4',
        'twitter_user4@email.com',
        '$2y$12$/HayGBMevM7EkbRho3xlC.nu/8.OeHGNi9iqGneNIDBrqnkD9fa1.',
        now()
    ) ON CONFLICT DO NOTHING
RETURNING id AS twitter_user4_id;