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
    status TEXT NOT NULL,
    card_token TEXT,
    payment_method_id TEXT,
    customer_id TEXT NOT NULL,
    phone TEXT NOT NULL
);

-- Table Children
CREATE TABLE IF NOT EXISTS children (
    id SERIAL,
    name VARCHAR(100) NOT NULL,
    rg VARCHAR(9) PRIMARY KEY NOT NULL,
    responsible_id VARCHAR(11) NOT NULL,
    shift TEXT NOT NULL,
    FOREIGN KEY (responsible_id) REFERENCES responsible(cpf) ON DELETE CASCADE
);