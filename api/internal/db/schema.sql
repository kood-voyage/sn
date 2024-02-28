CREATE TABLE
    IF NOT EXISTS user (
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

CREATE TABLE
    IF NOT EXISTS follower (
        id text PRIMARY KEY UNIQUE NOT NULL,
        source_id text NOT NULL,
        target_id text NOT NULL,
        FOREIGN KEY (source_id) REFERENCES user (id),
        FOREIGN KEY (target_id) REFERENCES user (id)
);

CREATE TABLE
    IF NOT EXISTS post (
        id text PRIMARY KEY UNIQUE NOT NULL,
        title text NOT NULL,
        content text NOT NULL,
        user_id text NOT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (user_id) REFERENCES user (id)
);

CREATE TABLE
    IF NOT EXISTS comment (
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

CREATE TABLE
    IF NOT EXISTS event (
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

CREATE TABLE
    IF NOT EXISTS community (
        id text PRIMARY KEY UNIQUE NOT NULL,
        creator_id text NOT NULL,
        name text NOT NULL,
        description text NOT NULL,
        FOREIGN KEY (creator_id) REFERENCES user (id)
);

CREATE TABLE
    IF NOT EXISTS event_option (
        id text PRIMARY KEY UNIQUE NOT NULL,
        event_id text NOT NULL,
        description text NOT NULL,
        FOREIGN KEY (event_id) REFERENCES event (id)
);

CREATE TABLE
    IF NOT EXISTS selected_event_option(
        id text PRIMARY KEY UNIQUE NOT NULL,
        option_id text NOT NULL,
        user_id text NOT NULL,
        FOREIGN KEY (option_id) REFERENCES event_option (id),
        FOREIGN KEY (user_id) REFERENCES user (id)
);

CREATE TABLE
    IF NOT EXISTS chat (
        id text PRIMARY KEY UNIQUE NOT NULL
);

CREATE TABLE
    IF NOT EXISTS chat_users (
        id text PRIMARY KEY UNIQUE NOT NULL,
        user_id text NOT NULL,
        chat_id text NOT NULL,
        FOREIGN KEY (user_id) REFERENCES user (id),
        FOREIGN KEY (chat_id) REFERENCES chat (id)
);

CREATE TABLE
    IF NOT EXISTS chat_lines (
        id text PRIMARY KEY UNIQUE NOT NULL,
        chat_id text NOT NULL,
        user_id text NOT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        message text NOT NULL
);

CREATE TABLE
    IF NOT EXISTS request (
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

CREATE TABLE
    IF NOT EXISTS request_type (
        id text PRIMARY KEY UNIQUE NOT NULL,
        description text NOT NULL
);

CREATE TABLE
    IF NOT EXISTS privacy (
        id text PRIMARY KEY UNIQUE NOT NULL,
        type_id text NOT NULL,
        FOREIGN KEY (type_id) REFERENCES privacy_type (id)
);

CREATE TABLE
    IF NOT EXISTS privacy_type (
        id text PRIMARY KEY UNIQUE NOT NULL,
        description text NOT NULL
);

CREATE TABLE
    IF NOT EXISTS reaction (
        id text PRIMARY KEY UNIQUE NOT NULL,
        emoji text NOT NULL,
        description text NOT NULL
);


CREATE TABLE
    IF NOT EXISTS reaction_join (
        id text PRIMARY KEY UNIQUE NOT NULL,
        user_id text NOT NULL,
        reaction_id text NOT NULL,
        parent_id text NOT NULL,
        FOREIGN KEY (user_id) REFERENCES user (id),
        FOREIGN KEY (reaction_id) REFERENCES reaction (id)
);

CREATE TABLE
    IF NOT EXISTS member (
        id text PRIMARY KEY UNIQUE NOT NULL,
        user_id text NOT NULL,
        group_id text NOT NULL,
        type_id text NOT NULL,
        FOREIGN KEY (user_id) REFERENCES user (id),
        FOREIGN KEY (group_id) REFERENCES community (id),
        FOREIGN KEY (type_id) REFERENCES member_type (id)
);

CREATE TABLE
    IF NOT EXISTS member_type (
        id text PRIMARY KEY UNIQUE NOT NULL,
        description text NOT NULL
);

CREATE TABLE
    IF NOT EXISTS image (
        id text PRIMARY KEY UNIQUE NOT NULL,
        parent_id text NOT NULL,
        type_id text NOT NULL,
        path text NOT NULL,
        FOREIGN KEY (type_id) REFERENCES image_type (id)
);

CREATE TABLE
    IF NOT EXISTS image_type (
        id text PRIMARY KEY UNIQUE NOT NULL,
        description text NOT NULL
);

CREATE TABLE
    IF NOT EXISTS session (
        id text PRIMARY KEY UNIQUE NOT NULL,
        user_id text NOT NULL,
        access_token_id text NOT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        expires_at DATETIME NOT NULL
);
