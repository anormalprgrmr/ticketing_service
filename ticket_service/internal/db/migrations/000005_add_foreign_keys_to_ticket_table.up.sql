ALTER TABLE tickets
ADD COLUMN support_id uuid,
ADD COLUMN user_id uuid;

ALTER TABLE tickets
ADD CONSTRAINT fk_ticket_support
FOREIGN KEY (support_id)
REFERENCES supports(id);

ALTER TABLE tickets
ADD CONSTRAINT fk_ticket_user
FOREIGN KEY (user_id)
REFERENCES users(id);