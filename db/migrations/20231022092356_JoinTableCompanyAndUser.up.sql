CREATE TABLE companies_spill_users (
    spill_user_id BIGINT REFERENCES spill_users(id) ON UPDATE CASCADE ON DELETE CASCADE,
    company_id BIGINT REFERENCES companies(id) ON UPDATE CASCADE,
    CONSTRAINT companies_spill_users_pkey PRIMARY KEY (spill_user_id, company_id)
)
