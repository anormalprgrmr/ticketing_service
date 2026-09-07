CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS tickets(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    body VARCHAR (300) NOT NULL
);
