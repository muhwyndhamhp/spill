ALTER TABLE spill_users 
ADD COLUMN created_at TIMESTAMPTZ NOT NULL, 
ADD COLUMN updated_at TIMESTAMPTZ NOT NULL,
ADD COLUMN deleted_at TIMESTAMPTZ;
