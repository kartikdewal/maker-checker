CREATE TABLE IF NOT EXISTS document_request (
    id              UUID         NOT NULL DEFAULT (uuid_generate_v4()),
    document_id     UUID         NOT NULL,
    creator_id      UUID         NOT NULL,
    approvers       JSONB,
    approver_count  INT          NOT NULL,
    status          VARCHAR(255) NOT NULL DEFAULT 'Pending',
    recipient_email VARCHAR(255) NOT NULL,
    created_at      TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP
);