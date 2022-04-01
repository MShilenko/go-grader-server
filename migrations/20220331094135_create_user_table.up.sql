CREATE TABLE IF NOT EXISTS "user" (
    id       INTEGER PRIMARY KEY NOT NULL,
    email VARCHAR(100) NOT NULL,
    password BYTEA NOT NULL,
    CONSTRAINT "user_email" UNIQUE ("email")
);
