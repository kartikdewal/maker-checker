package document

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"maker-checker/logger"
	"maker-checker/pkg/common"
	"maker-checker/pkg/document"
)

type Repository struct {
	log logger.ContextLogger
	db  *sqlx.DB
}

func NewRepository(log logger.ContextLogger, db *sqlx.DB) *Repository {
	return &Repository{log: log, db: db}
}

func (r *Repository) FindByID(ctx context.Context, id string) (*document.Row, error) {
	var result document.Row
	query := `SELECT * FROM document WHERE id = $1`
	err := r.db.GetContext(ctx, &result, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, common.ErrDocumentNotFound
		}

		return nil, err
	}

	return &result, nil
}

func (r *Repository) Create(ctx context.Context, input document.Row) (string, error) {
	query := `INSERT INTO document (id, description, creator_id, status, created_at, updated_at) VALUES (:id, :description, :creator_id, :status, :created_at, :updated_at)`
	_, err := r.db.NamedExecContext(ctx, query, input)
	if err != nil {
		r.log.Errorw(ctx, "failed to create document", "err", err)
		return "", err
	}

	return input.ID, err
}

func (r *Repository) Update(ctx context.Context, input document.Row) (string, error) {
	query := `UPDATE document SET description = :description, status = :status, updated_at = :updated_at WHERE id = :id`
	_, err := r.db.NamedExecContext(ctx, query, input)
	if err != nil {
		r.log.Errorw(ctx, "failed to update document", "err", err)
		return "", err
	}

	return input.ID, err
}
