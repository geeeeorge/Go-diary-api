DROP TABLE IF EXISTS memo;

CREATE TABLE memo (
    id SERIAL PRIMARY KEY,
    "date" VARCHAR(16) NOT NULL,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    "check" BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);
