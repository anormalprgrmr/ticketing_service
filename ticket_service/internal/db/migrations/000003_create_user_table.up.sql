CREATE TABLE IF NOT EXISTS users(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR (50) NOT NULL
);