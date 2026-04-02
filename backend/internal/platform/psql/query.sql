-- name: InsertAlert :exec
INSERT INTO alerts (fingerprint, status, starts_at, resolved_at, labels, annotations)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT (fingerprint, status, starts_at) DO NOTHING;
