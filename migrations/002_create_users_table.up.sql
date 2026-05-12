CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    email VARCHAR(255) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    dormitory_id varchar(2) NOT NULL REFERENCES dormitory (id) ON DELETE CASCADE,
    role VARCHAR(10) NOT NULL DEFAULT 'student',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_users_dormitory ON users (dormitory_id);

-- students 9 dormitory
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
        '4c6f1007-c7b8-49f7-a90a-d2df9ea5f402',
        'test9_2@mail.ru',
        '$2a$10$vVCHf1AuH6xsHCYduS1BYejGCRRjDSK/GzdUgSj9IMYawq4ANP1mq',
        '9'
    )
ON CONFLICT DO NOTHING;

-- admins 9 dormitory
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

-- students 1 dormitory
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
        'dc2e8989-6f7d-4f0d-978c-7ac764758536',
        'test1_2@mail.ru',
        '$2a$10$vVCHf1AuH6xsHCYduS1BYejGCRRjDSK/GzdUgSj9IMYawq4ANP1mq',
        '1',
    )
ON CONFLICT DO NOTHING;

-- admins 1 dormitory
INSERT INTO
    users (
        id,
        email,
        password,
        dormitory_id,
        role
    )
VALUES (
        '0b85c6d1-02cb-4d20-a777-b6d1f5aa922a',
        'admin1@mail.ru',
        '$2a$10$vVCHf1AuH6xsHCYduS1BYejGCRRjDSK/GzdUgSj9IMYawq4ANP1mq',
        '1',
        'admin'
    )
ON CONFLICT DO NOTHING;

-- students 7 dormitory
INSERT INTO
    users (
        id,
        email,
        password,
        dormitory_id
    )
VALUES (
        '62c258b3-42e0-45bc-8666-910657ea3890',
        'test7@mail.ru',
        '$2a$10$vVCHf1AuH6xsHCYduS1BYejGCRRjDSK/GzdUgSj9IMYawq4ANP1mq',
        '7'
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
        '638df0b7-4ebd-4e5b-a0e2-adc130317af5',
        'test7_2@mail.ru',
        '$2a$10$vVCHf1AuH6xsHCYduS1BYejGCRRjDSK/GzdUgSj9IMYawq4ANP1mq',
        '7',
    )
ON CONFLICT DO NOTHING;

-- admins 7 dormitory
INSERT INTO
    users (
        id,
        email,
        password,
        dormitory_id,
        role
    )
VALUES (
        '23d6a7f4-a9c2-49e5-ae36-fd8665c9f9e3',
        'admin7@mail.ru',
        '$2a$10$vVCHf1AuH6xsHCYduS1BYejGCRRjDSK/GzdUgSj9IMYawq4ANP1mq',
        '7',
        'admin'
    )
ON CONFLICT DO NOTHING;