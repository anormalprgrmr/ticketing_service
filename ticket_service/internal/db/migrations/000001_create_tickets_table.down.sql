CREATE TABLE IF NOT EXISTS tickets(
    id uuid DEFAULT gen_random_uuid(),
    support_id VARCHAR (50) NULL,
    password VARCHAR (50) NOT NULL,
    body VARCHAR (300) NOT NULL
    responses 
);