CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    email VARCHAR(255) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    dormitory_id varchar(2) NOT NULL REFERENCES dormitory (id) ON DELETE CASCADE,
    role VARCHAR(10) NOT NULL DEFAULT 'student',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_users_dormitory ON users (dormitory_id);