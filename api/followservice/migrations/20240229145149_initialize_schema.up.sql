CREATE TABLE follower (
                          source_id text NOT NULL,
                          target_id text NOT NULL
);

INSERT INTO follower (source_id, target_id) VALUES
                                                    ('user1', 'user2'),
                                                    ('user2', 'user3'),
                                                    ('user3', 'user4'),
                                                    ('user4', 'user5'),
                                                    ('user5', 'user1');

