CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    email VARCHAR(255) UNIQUE NOT NULL,
    dormitory_id UUID NOT NULL REFERENCES ON dormitory,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE dormitory (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid ()
);

CREATE INDEX idx_users_dormitory ON users (dormitory_id);