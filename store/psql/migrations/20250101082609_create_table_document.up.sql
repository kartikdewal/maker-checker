CREATE TABLE IF NOT EXISTS document (
    id           UUID         NOT NULL DEFAULT (uuid_generate_v4()),
    description  TEXT,
    creator_id   UUID         NOT NULL,
    status       VARCHAR(255) NOT NULL DEFAULT 'DRAFT',
    created_at   TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP
);