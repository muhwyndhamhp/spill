CREATE TABLE spill_users(
    id BIGSERIAL PRIMARY KEY,
    alias TEXT NOT NULL,
    email TEXT NOT NULL,
    bio TEXT
);

CREATE INDEX IF NOT EXISTS "idx_spill_users_id" ON spill_users(id);
