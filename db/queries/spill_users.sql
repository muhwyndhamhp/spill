-- name: GetSpillUsers :many
SELECT * FROM spill_users ORDER BY id;

-- name: GetSpillUserByID :one
SELECT * FROM spill_users WHERE id = $1 LIMIT 1;

-- name: GetSpillUserByServiceID :one
SELECT * FROM spill_users WHERE service_id = $1 LIMIT 1;

-- name: CreateSpillUser :one
INSERT INTO spill_users(alias, service_id, bio, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id;

-- name: GetSpillUsersFromCompanyID :many
SELECT u.* FROM companies_spill_users AS csu JOIN spill_users AS u ON csu.spill_user_id = u.id WHERE csu.company_id = $1 ORDER BY u.id;

-- name: SpillUserByAliasExist :one
SELECT EXISTS(SELECT * FROM spill_users WHERE alias = $1);

