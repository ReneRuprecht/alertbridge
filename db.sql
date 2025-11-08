CREATE TABLE IF NOT EXISTS alerts (
    id SERIAL PRIMARY KEY,
    alertname TEXT NOT NULL,
    status TEXT NOT NULL,
    labels JSONB,
    fingerprint TEXT UNIQUE
);
