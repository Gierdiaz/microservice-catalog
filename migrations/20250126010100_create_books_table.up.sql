CREATE EXTENSION IF NOT EXISTS "pgcrypto";

DROP TABLE IF EXISTS books;

CREATE TABLE books (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100),
    title VARCHAR(100),
    author VARCHAR(100),
    genre VARCHAR(100),
    price INT,
    quantity INT,
    year INT,
    available BOOLEAN,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
