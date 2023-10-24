-- name: CreateCompany :one
INSERT INTO companies(name, overall_score) VALUES($1, 0.0) RETURNING id;

-- name: GetCompanies :many
SELECT * FROM companies ORDER BY id;

-- name: GetCompanyByID :one
SELECT * FROM companies WHERE id = $1 LIMIT 1;

-- name: GetCompaniesByUserID :many
SELECT companies.* FROM companies_spill_users AS csu JOIN companies ON csu.company_id = companies.id WHERE csu.spill_user_id = $1 ORDER BY companies.id;

-- name: GetCompaniesByServiceID :many
SELECT c.* FROM companies_spill_users AS csu JOIN spill_users AS u ON csu.spill_user_id = u.id JOIN companies as c ON csu.company_id = c.id WHERE u.service_id = $1 ORDER BY c.id;

-- name: CompanyByServiceIDExist :one
SELECT EXISTS (SELECT c.id FROM companies_spill_users AS csu JOIN spill_users AS u ON csu.spill_user_id = u.id JOIN companies as c ON csu.company_id = c.id WHERE u.service_id = $1 ORDER BY c.id);

-- name: GetCompanyByName :one
SELECT * FROM companies WHERE LOWER(name) = $1;
