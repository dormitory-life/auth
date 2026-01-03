CREATE TABLE IF NOT EXISTS dormitory (
    id varchar(2) PRIMARY KEY,
    name TEXT NOT NULL,
    address TEXT NOT NULL,
    support_email TEXT DEFAULT '',
    description TEXT NOT NULL
);

INSERT INTO
    dormitory (
        id,
        name,
        address,
        description
    )
VALUES (
        '9',
        'Общежитие 9',
        'г Москва, ул Цимлянская, д 5',
        'Общежитие для иностранных студентов.'
    )
ON CONFLICT DO NOTHING;