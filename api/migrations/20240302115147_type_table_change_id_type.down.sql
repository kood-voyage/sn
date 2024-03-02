CREATE TABLE privacy_type_new (
                                  id text PRIMARY KEY UNIQUE NOT NULL,
                                  description TEXT NOT NULL
);

INSERT INTO privacy_type_new (id, description)
SELECT CAST(id AS TEXT), description FROM privacy_type;

DROP TABLE privacy_type;

ALTER TABLE privacy_type_new RENAME TO privacy_type;

CREATE TABLE member_type_new (
                                 id text PRIMARY KEY UNIQUE NOT NULL,
                                 description TEXT NOT NULL
);

INSERT INTO member_type_new (id, description)
SELECT CAST(id AS TEXT), description FROM member_type;

DROP TABLE member_type;

ALTER TABLE member_type_new RENAME TO member_type;

CREATE TABLE image_type_new (
                                id text PRIMARY KEY UNIQUE NOT NULL,
                                description TEXT NOT NULL
);

INSERT INTO image_type_new (id, description)
SELECT CAST(id AS TEXT), description FROM image_type;

DROP TABLE image_type;

ALTER TABLE image_type_new RENAME TO image_type;

CREATE TABLE request_type_new (
                                  id text PRIMARY KEY UNIQUE NOT NULL,
                                  description TEXT NOT NULL
);

INSERT INTO request_type_new (id, description)
SELECT CAST(id AS TEXT), description FROM request_type;

DROP TABLE request_type;

ALTER TABLE request_type_new RENAME TO request_type;