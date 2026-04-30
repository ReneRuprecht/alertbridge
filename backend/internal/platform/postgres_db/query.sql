-- name: InsertAlert :exec
INSERT INTO alerts (fingerprint, instance, status, starts_at, received_at, labels, annotations)
VALUES ($1, $2, $3, $4, $5, $6, $7)
ON CONFLICT (fingerprint, status, starts_at) DO NOTHING;

-- name: ListAlertsByInstance :many
SELECT fingerprint, instance, status, starts_at, received_at, labels, annotations FROM alerts 
WHERE instance=$1
ORDER BY received_at DESC;

-- name: InsertRule :exec
INSERT INTO rules (id, name, description, priority, enabled)
VALUES ($1, $2, $3, $4, $5);

-- name: ListRules :many
SELECT id, name, description, priority, enabled from rules;

-- name: InsertAction :exec
INSERT INTO actions (id, name, description, type, config)
VALUES ($1, $2, $3, $4, $5);

-- name: ListActions :many
SELECT id, name, description, type, config from actions;

-- name: FindActionById :one
SELECT id, name, description, type, config FROM actions 
WHERE id=$1
LIMIT 1;
