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
        'Москва, ул. Цимлянская, д. 5',
        'Общежитие для иностранных студентов коридорного типа.'
    )
ON CONFLICT DO NOTHING;

INSERT INTO
    dormitory (
        id,
        name,
        address,
        description
    )
VALUES (
        '1',
        'Общежитие 1',
        'Москва, ул. Б. Переяславская, д. 50, стр.1',
        'Общежитие для иностранных студентов блочного типа.'
    )
ON CONFLICT DO NOTHING;