CREATE TABLE IF NOT EXISTS users (
    user_id uuid DEFAULT gen_random_uuid() NOT NULL CONSTRAINT users_pk PRIMARY KEY,
    username varchar(50) NOT NULL,
    email varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    role varchar(20) NOT NULL,
    created_at timestamp DEFAULT NOW(),
    updated_at timestamp DEFAULT NOW(),
    deleted_at timestamp DEFAULT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS user_username_idx ON users (username)
WHERE
    (deleted_at IS NULL);

CREATE UNIQUE INDEX IF NOT EXISTS user_email_idx ON users (email, role)
WHERE
    (deleted_at IS NULL);