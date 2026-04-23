CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE alerts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    fingerprint TEXT NOT NULL,
    instance TEXT NOT NULL,
    status TEXT NOT NULL,
    starts_at TIMESTAMP WITH TIME ZONE NOT NULL,
    received_at TIMESTAMP WITH TIME ZONE NOT NULL,
    labels JSONB,
    annotations JSONB
);

ALTER TABLE alerts
ADD CONSTRAINT unique_alert UNIQUE(fingerprint, status, starts_at);

CREATE TABLE rules (
    id UUID PRIMARY KEY NOT NULL,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    priority INTEGER NOT NULL,
    enabled BOOLEAN NOT NULL
);
