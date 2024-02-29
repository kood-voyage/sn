CREATE TABLE user (
                      id text PRIMARY KEY UNIQUE NOT NULL,
                      username text UNIQUE NOT NULL,
                      email text UNIQUE NOT NULL,
                      password text NOT NULL,
                      created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                      date_of_birth text NOT NULL,
                      first_name text,
                      last_name text,
                      gender text,
                      description text
);

CREATE TABLE follower (
                          id text PRIMARY KEY UNIQUE NOT NULL,
                          source_id text NOT NULL,
                          target_id text NOT NULL,
                          FOREIGN KEY (source_id) REFERENCES user (id),
                          FOREIGN KEY (target_id) REFERENCES user (id)
);

CREATE TABLE post (
                      id text PRIMARY KEY UNIQUE NOT NULL,
                      title text NOT NULL,
                      content text NOT NULL,
                      user_id text NOT NULL,
                      created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                      FOREIGN KEY (user_id) REFERENCES user (id)
);

CREATE TABLE comment (
                         id text PRIMARY KEY UNIQUE NOT NULL,
                         user_id text NOT NULL,
                         post_id text NOT NULL,
                         parent_id text,
                         content text NOT NULL,
                         created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                         FOREIGN KEY (user_id) REFERENCES user (id),
                         FOREIGN KEY (post_id) REFERENCES post (id),
                         FOREIGN KEY (parent_id) REFERENCES comment (id)
);

CREATE TABLE event (
                       id text PRIMARY KEY UNIQUE NOT NULL,
                       user_id text NOT NULL,
                       group_id text NOT NULL,
                       name text NOT NULL,
                       description text NOT NULL,
                       created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                       date DATETIME,
                       FOREIGN KEY (user_id) REFERENCES user (id),
                       FOREIGN KEY (group_id) REFERENCES community (id)
);

CREATE TABLE community (
                           id text PRIMARY KEY UNIQUE NOT NULL,
                           creator_id text NOT NULL,
                           name text NOT NULL,
                           description text NOT NULL,
                           FOREIGN KEY (creator_id) REFERENCES user (id)
);

CREATE TABLE event_option (
                              id text PRIMARY KEY UNIQUE NOT NULL,
                              event_id text NOT NULL,
                              description text NOT NULL,
                              FOREIGN KEY (event_id) REFERENCES event (id)
);

CREATE TABLE selected_event_option(
                                      id text PRIMARY KEY UNIQUE NOT NULL,
                                      option_id text NOT NULL,
                                      user_id text NOT NULL,
                                      FOREIGN KEY (option_id) REFERENCES event_option (id),
                                      FOREIGN KEY (user_id) REFERENCES user (id)
);

CREATE TABLE chat (
                      id text PRIMARY KEY UNIQUE NOT NULL
);

CREATE TABLE chat_users (
                            id text PRIMARY KEY UNIQUE NOT NULL,
                            user_id text NOT NULL,
                            chat_id text NOT NULL,
                            FOREIGN KEY (user_id) REFERENCES user (id),
                            FOREIGN KEY (chat_id) REFERENCES chat (id)
);

CREATE TABLE chat_lines (
                            id text PRIMARY KEY UNIQUE NOT NULL,
                            chat_id text NOT NULL,
                            user_id text NOT NULL,
                            created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                            message text NOT NULL
);

CREATE TABLE request (
                         id text PRIMARY KEY UNIQUE NOT NULL,
                         type_id text NOT NULL,
                         source_id text NOT NULL,
                         target_id text NOT NULL,
                         message text NOT NULL,
                         created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                         FOREIGN KEY (type_id) REFERENCES request_type (id),
                         FOREIGN KEY (source_id) REFERENCES user (id),
                         FOREIGN KEY (target_id) REFERENCES user (id)
);

CREATE TABLE request_type (
                              id text PRIMARY KEY UNIQUE NOT NULL,
                              description text NOT NULL
);

CREATE TABLE privacy (
                         id text PRIMARY KEY UNIQUE NOT NULL,
                         type_id text NOT NULL,
                         FOREIGN KEY (type_id) REFERENCES privacy_type (id)
);

CREATE TABLE privacy_type (
                              id text PRIMARY KEY UNIQUE NOT NULL,
                              description text NOT NULL
);

CREATE TABLE reaction (
                          id text PRIMARY KEY UNIQUE NOT NULL,
                          emoji text NOT NULL,
                          description text NOT NULL
);


CREATE TABLE reaction_join (
                               id text PRIMARY KEY UNIQUE NOT NULL,
                               user_id text NOT NULL,
                               reaction_id text NOT NULL,
                               parent_id text NOT NULL,
                               FOREIGN KEY (user_id) REFERENCES user (id),
                               FOREIGN KEY (reaction_id) REFERENCES reaction (id)
);

CREATE TABLE member (
                        id text PRIMARY KEY UNIQUE NOT NULL,
                        user_id text NOT NULL,
                        group_id text NOT NULL,
                        type_id text NOT NULL,
                        FOREIGN KEY (user_id) REFERENCES user (id),
                        FOREIGN KEY (group_id) REFERENCES community (id),
                        FOREIGN KEY (type_id) REFERENCES member_type (id)
);

CREATE TABLE member_type (
                             id text PRIMARY KEY UNIQUE NOT NULL,
                             description text NOT NULL
);

CREATE TABLE image (
                       id text PRIMARY KEY UNIQUE NOT NULL,
                       parent_id text NOT NULL,
                       type_id text NOT NULL,
                       path text NOT NULL,
                       FOREIGN KEY (type_id) REFERENCES image_type (id)
);

CREATE TABLE image_type (
                            id text PRIMARY KEY UNIQUE NOT NULL,
                            description text NOT NULL
);

CREATE TABLE session (
                         id text PRIMARY KEY UNIQUE NOT NULL,
                         user_id text NOT NULL,
                         access_token_id text NOT NULL,
                         expires_at DATETIME NOT NULL
);
