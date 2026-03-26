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
        '4305c64b-fb4c-435b-968c-74cb49018702',
        'test9@mail.ru',
        '$2a$10$vVCHf1AuH6xsHCYduS1BYejGCRRjDSK/GzdUgSj9IMYawq4ANP1mq',
        '9'
    )
ON CONFLICT DO NOTHING;

INSERT INTO
    users (
        id,
        email,
        password,
        dormitory_id
    )
VALUES (
        '101f1810-5f46-4f06-94db-5b23b5695c2c',
        'test1@mail.ru',
        '$2a$10$vVCHf1AuH6xsHCYduS1BYejGCRRjDSK/GzdUgSj9IMYawq4ANP1mq',
        '1'
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
        'c01c19fd-f20f-4dbd-af11-4ba3ff0362b9',
        'admin9@mail.ru',
        '$2a$10$vVCHf1AuH6xsHCYduS1BYejGCRRjDSK/GzdUgSj9IMYawq4ANP1mq',
        '9',
        'admin'
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
        'dc2e8989-6f7d-4f0d-978c-7ac764758536',
        'admin1@mail.ru',
        '$2a$10$vVCHf1AuH6xsHCYduS1BYejGCRRjDSK/GzdUgSj9IMYawq4ANP1mq',
        '1',
        'admin'
    )
ON CONFLICT DO NOTHING;