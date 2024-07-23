-- Table Responsible

CREATE TABLE IF NOT EXISTS responsible (
    id SERIAL,
    name VARCHAR(100) NOT NULL,
    cpf VARCHAR(11) PRIMARY KEY NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    street VARCHAR(100) NOT NULL,
    number TEXT NOT NULL,
    complement TEXT,
    zip VARCHAR(8) NOT NULL,
    status TEXT NOT NULL
);