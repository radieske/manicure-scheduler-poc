CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE appointments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    phone TEXT NOT NULL,
    service TEXT NOT NULL,
    scheduled_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);