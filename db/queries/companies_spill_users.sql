-- name: CreateCompanySpillUser :exec
INSERT INTO companies_spill_users(spill_user_id, company_id) VALUES ($1, $2);
