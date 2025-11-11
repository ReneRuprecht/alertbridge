CREATE TABLE IF NOT EXISTS alerts_history (
    id SERIAL PRIMARY KEY,
    alertname TEXT NOT NULL,
    status TEXT NOT NULL,
    labels JSONB,
    fingerprint TEXT NOT NULL,
    starts_at TIMESTAMPTZ,
    ended_at TIMESTAMPTZ, 
    updated_at TIMESTAMPTZ, 
    UNIQUE(fingerprint, status, starts_at)
);
