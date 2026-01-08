CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    email VARCHAR(255) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    dormitory_id varchar(2) NOT NULL REFERENCES dormitory (id) ON DELETE CASCADE,
    role VARCHAR(10) NOT NULL DEFAULT 'student',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_users_dormitory ON users (dormitory_id);

INSERT INTO
    users (
        id,
        email,
        password,
        dormitory_id
    )
VALUES (
        'eeefdd1a-d422-480a-bbc8-e665e56dd76f',
        'test1@mail.ru',
        '$2a$10$vVCHf1AuH6xsHCYduS1BYejGCRRjDSK/GzdUgSj9IMYawq4ANP1mq',
        '9'
    )
ON CONFLICT DO NOTHING;

INSERT INTO
    users (
        id,
        email,
        password,
        dormitory_id,
        role
    )
VALUES (
        '3394a51a-8c0b-4ada-b6c2-a4e297d3aceb',
        'testAdmin@mail.ru',
        '$2a$10$vVCHf1AuH6xsHCYduS1BYejGCRRjDSK/GzdUgSj9IMYawq4ANP1mq',
        '9',
        'admin'
    )
ON CONFLICT DO NOTHING;