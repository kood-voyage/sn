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
                         parent_id text DEFAULT '',
                         message text NOT NULL,
                         created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                         FOREIGN KEY (type_id) REFERENCES request_type (id)
);
