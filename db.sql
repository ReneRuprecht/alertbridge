CREATE TABLE IF NOT EXISTS alerts_history (
    id SERIAL PRIMARY KEY,
    alertname TEXT NOT NULL,
    status TEXT NOT NULL,
    labels JSONB,
    fingerprint TEXT NOT NULL,
    starts_at TIMESTAMPTZ,
    ends_at TIMESTAMPTZ, 
    UNIQUE(fingerprint, status, labels, starts_at)
);
