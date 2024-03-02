CREATE TABLE privacy_type_new (
                                  id INTEGER PRIMARY KEY UNIQUE NOT NULL,
                                  description TEXT NOT NULL
);

INSERT INTO privacy_type_new (id, description)
SELECT CAST(id AS INTEGER), description FROM privacy_type;

INSERT INTO privacy_type_new (id, description) VALUES (0, 'public');
INSERT INTO privacy_type_new (id, description) VALUES (1, 'private');
INSERT INTO privacy_type_new (id, description) VALUES (2, 'selected');

DROP TABLE privacy_type;

ALTER TABLE privacy_type_new RENAME TO privacy_type;

CREATE TABLE member_type_new (
                                  id INTEGER PRIMARY KEY UNIQUE NOT NULL,
                                  description TEXT NOT NULL
);

INSERT INTO member_type_new (id, description)
SELECT CAST(id AS INTEGER), description FROM member_type;

DROP TABLE member_type;

ALTER TABLE member_type_new RENAME TO member_type;

CREATE TABLE image_type_new (
                                 id INTEGER PRIMARY KEY UNIQUE NOT NULL,
                                 description TEXT NOT NULL
);

INSERT INTO image_type_new (id, description)
SELECT CAST(id AS INTEGER), description FROM image_type;

DROP TABLE image_type;

ALTER TABLE image_type_new RENAME TO image_type;

CREATE TABLE request_type_new (
                                id INTEGER PRIMARY KEY UNIQUE NOT NULL,
                                description TEXT NOT NULL
);

INSERT INTO request_type_new (id, description)
SELECT CAST(id AS INTEGER), description FROM request_type;

DROP TABLE request_type;

ALTER TABLE request_type_new RENAME TO request_type;