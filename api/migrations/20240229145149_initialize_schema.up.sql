CREATE TABLE user (
                      id text PRIMARY KEY UNIQUE NOT NULL
);

INSERT INTO user (id) VALUES
('user1'),
('user2'),
('user3'),
('user4'),
('user5'),
('user6'),
('user7'),
('user8'),
('user9'),
('user10');

CREATE TABLE community (
                           id text PRIMARY KEY UNIQUE NOT NULL,
                           creator_id text NOT NULL,
                           name text NOT NULL,
                           description text NOT NULL,
                           FOREIGN KEY (creator_id) REFERENCES user (id)
);

INSERT INTO community (id, creator_id, name, description) VALUES
  ('group1', 'user1', 'Tech Enthusiasts', 'A community for tech lovers'),
  ('group2', 'user2', 'Book Club', 'Discussing and sharing thoughts on books'),
  ('group3', 'user3', 'Fitness Fanatics', 'For fitness and workout enthusiasts'),
  ('group4', 'user4', 'Foodies Corner', 'Exploring and sharing culinary experiences'),
  ('group5', 'user5', 'Travel Explorers', 'Adventures and travel experiences');

CREATE TABLE follower (
                          id text PRIMARY KEY UNIQUE NOT NULL,
                          source_id text NOT NULL,
                          target_id text NOT NULL,
                          FOREIGN KEY (source_id) REFERENCES user (id) ON DELETE CASCADE ,
                          FOREIGN KEY (target_id) REFERENCES user (id) ON DELETE CASCADE
);

INSERT INTO follower (id, source_id, target_id) VALUES
  ('follow1', 'user1', 'user2'),
  ('follow2', 'user2', 'user3'),
  ('follow3', 'user3', 'user4'),
  ('follow4', 'user4', 'user5'),
  ('follow5', 'user5', 'user1');

CREATE TABLE post (
                      id text PRIMARY KEY UNIQUE NOT NULL,
                      title text NOT NULL,
                      content text NOT NULL,
                      user_id text NOT NULL,
                      created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                      FOREIGN KEY (user_id) REFERENCES user (id)
);

INSERT INTO post (id, title, content, user_id, created_at) VALUES
  ('post1', 'Title 1', 'Content 1', 'user1', '2022-01-01 12:00:00'),
  ('post2', 'Title 2', 'Content 2', 'user2', '2022-01-02 13:30:00'),
  ('post3', 'Title 3', 'Content 3', 'user3', '2022-01-03 15:45:00'),
  ('post4', 'Title 4', 'Content 4', 'user4', '2022-01-04 09:15:00'),
  ('post5', 'Title 5', 'Content 5', 'user5', '2022-01-05 17:00:00');

CREATE TABLE comment (
                         id text PRIMARY KEY UNIQUE NOT NULL,
                         user_id text NOT NULL,
                         post_id text NOT NULL,
                         parent_id text DEFAULT NULL,
                         content text NOT NULL,
                         created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                         FOREIGN KEY (user_id) REFERENCES user (id) ON DELETE CASCADE ,
                         FOREIGN KEY (post_id) REFERENCES post (id) ON DELETE CASCADE ,
                         FOREIGN KEY (parent_id) REFERENCES comment (id) ON DELETE CASCADE
);

INSERT INTO comment (id, user_id, post_id, parent_id, content, created_at) VALUES
  ('comment1', 'user2', 'post1', NULL, 'Comment on post 1', '2022-01-01 12:30:00'),
  ('comment2', 'user3', 'post1', NULL, 'Another comment on post 1', '2022-01-01 13:00:00'),
  ('comment3', 'user4', 'post2', NULL, 'Comment on post 2', '2022-01-02 14:00:00'),
  ('comment4', 'user1', 'post3', NULL, 'Comment on post 3', '2022-01-03 16:00:00'),
  ('comment5', 'user5', 'post4', NULL, 'Comment on post 4', '2022-01-04 10:00:00'),
  ('comment6', 'user2', 'post5', NULL, 'Comment on post 5', '2022-01-05 18:00:00'),
  ('comment7', 'user3', 'post5', 'comment6', 'Reply to comment 6', '2022-01-05 18:30:00'),
  ('comment8', 'user4', 'post5', NULL, 'Another comment on post 5', '2022-01-05 19:00:00');

CREATE TABLE event (
                       id text PRIMARY KEY UNIQUE NOT NULL,
                       user_id text NOT NULL,
                       group_id text NOT NULL,
                       name text NOT NULL,
                       description text NOT NULL,
                       created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                       date DATETIME,
                       FOREIGN KEY (user_id) REFERENCES user (id),
                       FOREIGN KEY (group_id) REFERENCES community (id) ON DELETE CASCADE
);

INSERT INTO event (id, user_id, group_id, name, description, date)
VALUES
  ('event1', 'user1', 'group1', 'Event 1', 'Description for Event 1', '2024-03-01 12:00:00'),
  ('event2', 'user2', 'group1', 'Event 2', 'Description for Event 2', '2024-03-15 15:30:00');


CREATE TABLE event_option_type (
                              id INT PRIMARY KEY UNIQUE NOT NULL,
                              description text NOT NULL
);

INSERT INTO event_option_type (id, description) VALUES
  (1, 'Going'),
  (2, 'Not Going'),
  (3, 'Interested'),
  (4, 'Maybe');

CREATE TABLE event_registered_users (
                                      id text PRIMARY KEY UNIQUE NOT NULL,
                                      type_id INT NOT NULL,
                                      user_id text NOT NULL,
                                      event_id text NOT NULL,
                                      FOREIGN KEY (type_id) REFERENCES event_option_type (id),
                                      FOREIGN KEY (user_id) REFERENCES user (id) ON DELETE CASCADE,
                                      FOREIGN KEY (event_id) REFERENCES event (id) ON DELETE CASCADE
);

INSERT INTO event_registered_users (id, type_id, user_id, event_id)
VALUES
  ('registration1_id', 1, 'user3', 'event1'), -- Going
  ('registration2_id', 2, 'user4', 'event1'), -- Not Going
  ('registration3_id', 1, 'user5', 'event2'); -- Going

CREATE TABLE chat (
                      id text PRIMARY KEY UNIQUE NOT NULL
);

CREATE TABLE chat_users (
                            id text PRIMARY KEY UNIQUE NOT NULL,
                            user_id text NOT NULL,
                            chat_id text NOT NULL,
                            FOREIGN KEY (user_id) REFERENCES user (id) ON DELETE CASCADE ,
                            FOREIGN KEY (chat_id) REFERENCES chat (id) ON DELETE CASCADE
);

CREATE TABLE chat_lines (
                            id text PRIMARY KEY UNIQUE NOT NULL,
                            chat_id text NOT NULL,
                            user_id text NOT NULL,
                            created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                            message text NOT NULL
);

CREATE TABLE request_type (
                              id INT PRIMARY KEY UNIQUE NOT NULL,
                              description text NOT NULL
);

INSERT INTO request_type (id, description) VALUES (1, 'notification');
INSERT INTO request_type (id, description) VALUES (2, 'follow');
INSERT INTO request_type (id, description) VALUES (3, 'invite');


CREATE TABLE request (
                         id text PRIMARY KEY UNIQUE NOT NULL,
                         type_id INT NOT NULL,
                         source_id text NOT NULL,
                         target_id text NOT NULL,
                         message text NOT NULL,
                         created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                         FOREIGN KEY (type_id) REFERENCES request_type (id) ,
                         FOREIGN KEY (source_id) REFERENCES user (id) ON DELETE CASCADE ,
                         FOREIGN KEY (target_id) REFERENCES user (id) ON DELETE CASCADE
);


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


CREATE TABLE reaction_join (
                               id text PRIMARY KEY UNIQUE NOT NULL,
                               user_id text NOT NULL,
                               reaction_id text NOT NULL,
                               parent_id text NOT NULL,
                               FOREIGN KEY (user_id) REFERENCES user (id) ON DELETE CASCADE ,
                               FOREIGN KEY (reaction_id) REFERENCES reaction (id)
);

CREATE TABLE reaction (
                          id text PRIMARY KEY UNIQUE NOT NULL,
                          emoji text NOT NULL,
                          description text NOT NULL
);



CREATE TABLE member_type (
                             id INT PRIMARY KEY UNIQUE NOT NULL,
                             description text NOT NULL
);

INSERT INTO member_type (id, description) VALUES (1, 'admin');
INSERT INTO member_type (id, description) VALUES (2, 'user');

CREATE TABLE member (
                        id text PRIMARY KEY UNIQUE NOT NULL,
                        user_id text NOT NULL,
                        group_id text NOT NULL,
                        type_id INT NOT NULL,
                        FOREIGN KEY (user_id) REFERENCES user (id) ON DELETE CASCADE ,
                        FOREIGN KEY (group_id) REFERENCES community (id) ON DELETE CASCADE ,
                        FOREIGN KEY (type_id) REFERENCES member_type (id)
);

INSERT INTO member (id, user_id, group_id, type_id) VALUES
  ('member1', 'user1', 'group1', 1), -- Admin
  ('member2', 'user2', 'group1', 2), -- User
  ('member3', 'user3', 'group2', 1), -- Admin
  ('member4', 'user4', 'group2', 2), -- User
  ('member5', 'user5', 'group3', 1), -- Admin
  ('member6', 'user6', 'group3', 2), -- User
  ('member7', 'user7', 'group4', 1), -- Admin
  ('member8', 'user8', 'group4', 2), -- User
  ('member9', 'user9', 'group5', 1), -- Admin
  ('member10', 'user10', 'group5', 2); -- User


CREATE TABLE image_type (
                            id INT PRIMARY KEY UNIQUE NOT NULL,
                            description text NOT NULL
);
CREATE TABLE image (
                       id text PRIMARY KEY UNIQUE NOT NULL,
                       parent_id text NOT NULL,
                       type_id INT NOT NULL,
                       path text NOT NULL,
                       FOREIGN KEY (type_id) REFERENCES image_type (id)
);


INSERT INTO image_type (id, description) VALUES (1, 'banner');
INSERT INTO image_type (id, description) VALUES (2, 'avatar');
INSERT INTO image_type (id, description) VALUES (3, 'profile');
INSERT INTO image_type (id, description) VALUES (4, 'header');

CREATE TABLE selected_users (
    id text NOT NULL UNIQUE,
    user_id text NOT NULL,
    parent_id text NOT NULL,
    FOREIGN KEY (parent_id) REFERENCES privacy (id) ON DELETE CASCADE
);

INSERT INTO selected_users (id, user_id, parent_id)
VALUES
  ('selected_user1', 'user1', 'post2'),
  ('selected_user2', 'user3', 'post2'),
  ('selected_user3', 'user4', 'post2'),
  ('selected_user5', 'user5', 'post5'),
  ('selected_user6', 'user2', 'post2'),
  ('selected_user7', 'user3', 'post5'),
  ('selected_user8', 'user4', 'post5');