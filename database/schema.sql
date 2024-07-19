-- Table Responsible
CREATE TABLE IF NOT EXISTS responsible (
    id SERIAL,
    name VARCHAR(100) NOT NULL,
    cpf VARCHAR(14) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    street VARCHAR(100) NOT NULL,
    number VARCHAR(10) NOT NULL,
    complement VARCHAR(10),
    zip VARCHAR(8) NOT NULL
);

-- Table Children
CREATE TABLE IF NOT EXISTS children (
    id SERIAL,
    name VARCHAR(100) NOT NULL,
    rg VARCHAR(9) PRIMARY KEY NOT NULL,
    responsible_id VARCHAR(14) NOT NULL,
    FOREIGN KEY (responsible_id) REFERENCES responsible(cpf) ON DELETE CASCADE
);