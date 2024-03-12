CREATE TABLE
  IF NOT EXISTS user (
    id text PRIMARY KEY UNIQUE NOT NULL,
    username text UNIQUE NOT NULL,
    email text UNIQUE NOT NULL,
    password text NOT NULL,
    timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
    date_of_birth text NOT NULL,
    first_name text NOT NULL,
    last_name text NOT NULL,
    description text
  );

CREATE TABLE
  IF NOT EXISTS session (
    access_id text PRIMARY KEY UNIQUE NOT NULL,
    user_id text UNIQUE NOT NULL,
    timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES user (id) ON DELETE CASCADE
  )