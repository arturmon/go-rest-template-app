CREATE TABLE IF NOT EXISTS users (
                                      _id      TEXT,
                                      name     TEXT,
                                      email    TEXT UNIQUE,
                                      password BYTEA
);

CREATE INDEX idx_user_email ON users (email);
alter table users owner to root;
