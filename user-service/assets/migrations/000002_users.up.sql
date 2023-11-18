
CREATE TABLE IF NOT EXISTS users (
    id TEXT PRIMARY KEY,

    created_at TEXT NOT NULL,
    updated_at TEXT NOT NULL,

    nickname TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,

    email       TEXT NOT NULL UNIQUE,
    is_verified INTEGER NOT NULL DEFAULT 0
);
