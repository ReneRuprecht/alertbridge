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

CREATE TABLE actions (
    id UUID PRIMARY KEY NOT NULL,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    type TEXT NOT NULL,
    config JSONB NOT NULL
);

CREATE TABLE rule_conditions (
    id UUID PRIMARY KEY NOT NULL,
    rule_id UUID REFERENCES rules(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    operator TEXT NOT NULL,
    field TEXT NOT NULL,
    value TEXT NOT NULL
);
