CREATE TABLE
  IF NOT EXISTS user (
    id text PRIMARY KEY UNIQUE NOT NULL,
    username text UNIQUE NOT NULL,
    email text UNIQUE NOT NULL,
    password text NOT NULL,
    timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
    date_of_birth text NOT NULL,
    first_name text,
    last_name text,
    gender text
  );