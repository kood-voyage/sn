CREATE TABLE
    user (
        id text PRIMARY KEY UNIQUE NOT NULL,
        username text UNIQUE NOT NULL,
        email text UNIQUE NOT NULL,
        password text NOT NULL,
        timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
        date_of_birth text NOT NULL,
        first_name text NOT NULL,
        last_name text NOT NULL,
        description text,
        avatar text,
        cover text
    );

CREATE TABLE
    session (
        access_id text PRIMARY KEY UNIQUE NOT NULL,
        user_id text UNIQUE NOT NULL,
        timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (user_id) REFERENCES user (id) ON DELETE CASCADE
    );

CREATE TABLE
    community (
        id text PRIMARY KEY UNIQUE NOT NULL,
        creator_id text NOT NULL,
        name text UNIQUE NOT NULL,
        description text NOT NULL,
        FOREIGN KEY (creator_id) REFERENCES user (id)
    );

CREATE TABLE
    follower (
        id text PRIMARY KEY UNIQUE NOT NULL,
        source_id text NOT NULL,
        target_id text NOT NULL,
        FOREIGN KEY (source_id) REFERENCES user (id) ON DELETE CASCADE,
        FOREIGN KEY (target_id) REFERENCES user (id) ON DELETE CASCADE
    );

CREATE TABLE
    post (
        id text PRIMARY KEY UNIQUE NOT NULL,
        title text NOT NULL,
        content text NOT NULL,
        user_id text NOT NULL,
        community_id text DEFAULT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (user_id) REFERENCES user (id)
    );

CREATE TABLE
    comment (
        id text PRIMARY KEY UNIQUE NOT NULL,
        user_id text NOT NULL,
        post_id text NOT NULL,
        parent_id text DEFAULT NULL,
        content text NOT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        user_name text NOT NULL,
        user_avatar text NOT NULL,
        FOREIGN KEY (user_id) REFERENCES user (id) ON DELETE CASCADE,
        FOREIGN KEY (post_id) REFERENCES post (id) ON DELETE CASCADE,
        FOREIGN KEY (parent_id) REFERENCES comment (id) ON DELETE CASCADE
    );

CREATE TABLE
    event (
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

CREATE TABLE
    event_option_type (
        id INT PRIMARY KEY UNIQUE NOT NULL,
        description text NOT NULL
    );

INSERT INTO
    event_option_type (id, description)
VALUES
    (1, 'Going'),
    (2, 'Not Going'),
    (3, 'Interested'),
    (4, 'Maybe');

CREATE TABLE
    event_registered_users (
        id text PRIMARY KEY UNIQUE NOT NULL,
        type_id INT NOT NULL,
        user_id text NOT NULL,
        event_id text NOT NULL,
        FOREIGN KEY (type_id) REFERENCES event_option_type (id),
        FOREIGN KEY (user_id) REFERENCES user (id) ON DELETE CASCADE,
        FOREIGN KEY (event_id) REFERENCES event (id) ON DELETE CASCADE
    );

CREATE TABLE
    chat (
        id text PRIMARY KEY UNIQUE NOT NULL,
        group_id text
    );

CREATE TABLE
    chat_users (
        id text PRIMARY KEY UNIQUE NOT NULL,
        user_id text NOT NULL,
        chat_id text NOT NULL,
        FOREIGN KEY (user_id) REFERENCES user (id) ON DELETE CASCADE,
        FOREIGN KEY (chat_id) REFERENCES chat (id) ON DELETE CASCADE
    );

CREATE TABLE
    chat_lines (
        id text PRIMARY KEY UNIQUE NOT NULL,
        chat_id text NOT NULL,
        user_id text NOT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        message text NOT NULL
    );

CREATE TABLE
    request_type (
        id INT PRIMARY KEY UNIQUE NOT NULL,
        description text NOT NULL
    );

INSERT INTO
    request_type (id, description)
VALUES
    (1, 'notification');

INSERT INTO
    request_type (id, description)
VALUES
    (2, 'follow');

INSERT INTO
    request_type (id, description)
VALUES
    (3, 'invite');

CREATE TABLE
    request (
        id text PRIMARY KEY UNIQUE NOT NULL,
        type_id INT NOT NULL,
        source_id text NOT NULL,
        target_id text NOT NULL,
        parent_id text DEFAULT '',
        message text NOT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (type_id) REFERENCES request_type (id),
        FOREIGN KEY (source_id) REFERENCES user (id) ON DELETE CASCADE,
        FOREIGN KEY (target_id) REFERENCES user (id) ON DELETE CASCADE
    );

CREATE TABLE
    privacy_type (
        id INT PRIMARY KEY UNIQUE NOT NULL,
        description text NOT NULL
    );

INSERT INTO
    privacy_type (id, description)
VALUES
    (1, 'public');

INSERT INTO
    privacy_type (id, description)
VALUES
    (2, 'private');

INSERT INTO
    privacy_type (id, description)
VALUES
    (3, 'selected');

CREATE TABLE
    privacy (
        id text PRIMARY KEY UNIQUE NOT NULL,
        type_id INT NOT NULL,
        FOREIGN KEY (type_id) REFERENCES privacy_type (id)
    );

CREATE TABLE
    reaction_join (
        id text PRIMARY KEY UNIQUE NOT NULL,
        user_id text NOT NULL,
        reaction_id text NOT NULL,
        parent_id text NOT NULL,
        FOREIGN KEY (user_id) REFERENCES user (id) ON DELETE CASCADE,
        FOREIGN KEY (reaction_id) REFERENCES reaction (id)
    );

CREATE TABLE
    reaction (
        id text PRIMARY KEY UNIQUE NOT NULL,
        emoji text NOT NULL,
        description text NOT NULL
    );

CREATE TABLE
    member_type (
        id INT PRIMARY KEY UNIQUE NOT NULL,
        description text NOT NULL
    );

INSERT INTO
    member_type (id, description)
VALUES
    (1, 'admin');

INSERT INTO
    member_type (id, description)
VALUES
    (2, 'user');

CREATE TABLE
    member (
        id text PRIMARY KEY UNIQUE NOT NULL,
        user_id text NOT NULL,
        group_id text NOT NULL,
        type_id INT NOT NULL,
        FOREIGN KEY (user_id) REFERENCES user (id) ON DELETE CASCADE,
        FOREIGN KEY (group_id) REFERENCES community (id) ON DELETE CASCADE,
        FOREIGN KEY (type_id) REFERENCES member_type (id)
    );

CREATE TABLE
    selected_users (
        id text NOT NULL UNIQUE,
        user_id text NOT NULL,
        parent_id text NOT NULL,
        FOREIGN KEY (parent_id) REFERENCES privacy (id) ON DELETE CASCADE
    );

CREATE TABLE
    image (
        id text NOT NULL UNIQUE,
        parent_id text NOT NULL,
        path text NOT NULL
    )