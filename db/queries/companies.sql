-- name: CreateCompany :one
INSERT INTO companies(id, name, overall_score) VALUES($1, $2, 0.0) RETURNING id;

-- name: GetCompanies :many
SELECT * FROM companies ORDER BY id;

-- name: GetCompanyByID :one
SELECT * FROM companies WHERE id = $1 LIMIT 1;

-- name: GetCompaniesByUserID :many
SELECT companies.* FROM companies_spill_users AS csu JOIN companies ON csu.company_id = companies.id WHERE csu.spill_user_id = $1 ORDER BY companies.id;
