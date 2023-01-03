CREATE TABLE IF NOT EXISTS users
(
    id         UUID PRIMARY KEY,
    first_name TEXT                     NOT NULL,
    last_name  TEXT                     NOT NULL,
    age        INT check ( age > 0 )    NOT NULL,
    gender     TEXT                     NOT NULL,
    biography  TEXT                     NOT NULL,
    city       TEXT                     NOT NULL,
    password   TEXT                     NOT NULL,
    created    TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);