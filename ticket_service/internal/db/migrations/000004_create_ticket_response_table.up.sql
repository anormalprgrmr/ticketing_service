CREATE TABLE IF NOT EXISTS ticket_responses(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    ticket_id uuid,
    created_at TIMESTAMP DEFAULT NOW(),
    body VARCHAR (300) NOT NULL,

    CONSTRAINT fk_response_ticket
        FOREIGN KEY (ticket_id)
        REFERENCES tickets(id)
        ON DELETE CASCADE
);