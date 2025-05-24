CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    is_activated BOOLEAN NOT NULL DEFAULT FALSE,

    confirmation_token TEXT,
    confirmation_expires_at TIMESTAMPTZ,
    confirmed_at TIMESTAMPTZ,

    reset_password_token TEXT,
    reset_password_expires_at TIMESTAMPTZ,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);