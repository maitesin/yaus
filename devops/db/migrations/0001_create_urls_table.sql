CREATE TABLE IF NOT EXISTS urls (
    original TEXT PRIMARY KEY,
    shortened TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL
);
CREATE UNIQUE INDEX IF NOT EXISTS idx_urls_shortened ON PUBLIC.urls(shortened);
CREATE INDEX IF NOT EXISTS idx_urls_createdat ON PUBLIC.urls(created_at);
