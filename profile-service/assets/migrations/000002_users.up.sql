BEGIN;

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    nickname VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,

    email       VARCHAR(255) NOT NULL UNIQUE,
    is_verified BOOLEAN NOT NULL DEFAULT false
);

COMMIT;
