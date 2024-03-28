CREATE TABLE privacy_type (
                              id INT PRIMARY KEY UNIQUE NOT NULL,
                              description text NOT NULL
);

INSERT INTO privacy_type (id, description) VALUES (1, 'public');
INSERT INTO privacy_type (id, description) VALUES (2, 'private');
INSERT INTO privacy_type (id, description) VALUES (3, 'selected');

CREATE TABLE privacy (
                         id text PRIMARY KEY UNIQUE NOT NULL,
                         type_id INT NOT NULL,
                         FOREIGN KEY (type_id) REFERENCES privacy_type (id)
);

INSERT INTO privacy (id, type_id)
VALUES
    ('user1', 1), -- Public
    ('user2', 2), -- Private
    ('user3', 1), -- Public
    ('user4', 2), -- Private
    ('user5', 3), -- Selected
    ('user6', 1), -- Public
    ('user7', 2), -- Private
    ('user8', 3), -- Selected
    ('user9', 1), -- Public
    ('post1', 2), -- Private
    ('post2', 3), -- Selected
    ('post3', 1), -- Public
    ('post4', 2), -- Private
    ('post5', 3), -- Selected
    ('group1', 2), -- Private
    ('group2', 1), -- Public
    ('group3', 2), -- Private
    ('group4', 1), -- Public
    ('group5', 2); -- Private