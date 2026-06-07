-- +goose Up
CREATE TABLE IF NOT EXISTS t_user (
    id UUID PRIMARY KEY,
    email TEXT NOT NULL,
    password_hash TEXT NOT NULL,
    nickname TEXT NOT NULL,
    role TEXT NOT NULL DEFAULT 'student',
    status TEXT NOT NULL DEFAULT 'active',
    last_login_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,
    CONSTRAINT t_user_email_unique UNIQUE (email),
    CONSTRAINT t_user_role_check CHECK (role IN ('student', 'teacher', 'admin')),
    CONSTRAINT t_user_status_check CHECK (status IN ('active', 'disabled')),
    CONSTRAINT t_user_email_not_blank CHECK (length(trim(email)) > 0),
    CONSTRAINT t_user_nickname_not_blank CHECK (length(trim(nickname)) > 0)
);

CREATE INDEX IF NOT EXISTS idx_t_user_role ON t_user (role);
CREATE INDEX IF NOT EXISTS idx_t_user_status ON t_user (status);
CREATE INDEX IF NOT EXISTS idx_t_user_deleted_at ON t_user (deleted_at);

-- +goose Down
DROP TABLE IF EXISTS t_user;
