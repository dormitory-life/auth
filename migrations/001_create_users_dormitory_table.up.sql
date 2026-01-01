CREATE TABLE dormitory (id INTEGER PRIMARY KEY);

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    email VARCHAR(255) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    dormitory_id INTEGER NOT NULL REFERENCES dormitory (id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_dormitory ON users (dormitory_id);

INSERT INTO
    dormitory (id)
VALUES (1),
    (2),
    (3),
    (4),
    (5),
    (6),
    (7),
    (8),
    (9)