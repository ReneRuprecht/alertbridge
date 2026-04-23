-- name: InsertAlert :exec
INSERT INTO alerts (fingerprint, instance, status, starts_at, received_at, labels, annotations)
VALUES ($1, $2, $3, $4, $5, $6, $7)
ON CONFLICT (fingerprint, status, starts_at) DO NOTHING;

-- name: FindAlertsByInstance :many
SELECT fingerprint, instance, status, starts_at, received_at, labels, annotations FROM alerts 
WHERE instance=$1
ORDER BY received_at DESC;

-- name: InsertRule :exec
INSERT INTO rules (id, name, description, priority, enabled)
VALUES ($1, $2, $3, $4, $5);

-- name: ListRules :many
SELECT id, name, description, priority, enabled from rules;
