ALTER TABLE supports
    ALTER COLUMN id SET DEFAULT uuidv7();

ALTER TABLE ticket_responses
    ALTER COLUMN id SET DEFAULT uuidv7();

ALTER TABLE tickets
    ALTER COLUMN id SET DEFAULT uuidv7();

ALTER TABLE users
    ALTER COLUMN id SET DEFAULT uuidv7();
