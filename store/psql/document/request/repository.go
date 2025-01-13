package request

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"maker-checker/logger"
	"maker-checker/pkg/common"
	"maker-checker/pkg/document/request"
)

type Repository struct {
	log logger.ContextLogger
	db  *sqlx.DB
}

func NewRepository(log logger.ContextLogger, db *sqlx.DB) *Repository {
	return &Repository{log: log, db: db}
}

func (r *Repository) FindByID(ctx context.Context, id string) (*request.Row, error) {
	var result request.Row
	query := `SELECT * FROM document_request WHERE id = $1`
	err := r.db.GetContext(ctx, &result, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, common.ErrDocumentRequestNotFound
		}

		return nil, err
	}
	err = result.Unmarshal()
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *Repository) Create(ctx context.Context, payload request.Row) (string, error) {
	payloadMap, err := payload.Marshal()
	if err != nil {
		r.log.Errorw(ctx, "error parsing document request", "err", err)
	}

	query := `INSERT INTO document_request (id, document_id, creator_id, approvers, approver_count, status, recipient_email) VALUES (:id, :document_id, :creator_id, :approvers, :approver_count, :status, :recipient_email)`
	_, err = r.db.NamedExecContext(ctx, query, payloadMap)
	if err != nil {
		r.log.Errorw(ctx, "failed to create document request", "err", err)
		return "", err
	}

	return payload.ID, err
}

func (r *Repository) Update(ctx context.Context, input request.Row) (string, error) {
	args, err := input.Marshal()
	if err != nil {
		r.log.Errorw(ctx, "error parsing document request", "err", err)
	}

	query := `UPDATE document_request SET status = :status, approvers = :approvers, updated_at = :updated_at WHERE id = :id`
	_, err = r.db.NamedExecContext(ctx, query, args)
	if err != nil {
		r.log.Errorw(ctx, "failed to update document request", "err", err)
		return "", err
	}

	return input.ID, err
}
