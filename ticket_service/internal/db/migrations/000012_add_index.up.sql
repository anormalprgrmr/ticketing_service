CREATE INDEX idx_ticket_responses_ticket_id
ON ticket_responses (ticket_id);

CREATE INDEX idx_tickets_user_id
ON tickets (user_id);

CREATE INDEX idx_tickets_support_id
ON tickets (support_id);