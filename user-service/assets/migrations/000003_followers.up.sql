BEGIN;

CREATE TABLE IF NOT EXISTS followers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    from_user_id UUID NOT NULL REFERENCES users (id),
    to_user_id   UUID NOT NULL REFERENCES users (id),

    UNIQUE (from_user_id, to_user_id)
);

COMMIT;
