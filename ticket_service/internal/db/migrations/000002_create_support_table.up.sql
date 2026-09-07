CREATE TABLE IF NOT EXISTS supports (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    current_assigned_ticket_id uuid,
    last_assigned_ticket_time TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_support_current_ticket
        FOREIGN KEY (current_assigned_ticket_id)
        REFERENCES tickets(id)
);