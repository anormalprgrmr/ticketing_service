CREATE TYPE ticket_status AS ENUM('Opened','Answered','Closed');

ALTER TABLE tickets
ADD COLUMN status ticket_status NOT NULL DEFAULT 'Opened';