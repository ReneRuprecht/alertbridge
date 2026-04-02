CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE alerts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    fingerprint TEXT NOT NULL,
    status TEXT NOT NULL,
    starts_at TIMESTAMP WITH TIME ZONE NOT NULL,
    resolved_at TIMESTAMP WITH TIME ZONE,
    labels JSONB,
    annotations JSONB
);

ALTER TABLE alerts
ADD CONSTRAINT unique_alert UNIQUE(fingerprint, status, starts_at);
