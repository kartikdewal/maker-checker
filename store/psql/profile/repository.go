package profile

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"maker-checker/logger"
	"maker-checker/pkg/common"
	"maker-checker/pkg/profile"
)

type Repository struct {
	log logger.ContextLogger
	db  *sqlx.DB
}

func NewRepository(log logger.ContextLogger, db *sqlx.DB) *Repository {
	return &Repository{log: log, db: db}
}

func (r *Repository) FindByID(ctx context.Context, profileID string) (*profile.Row, error) {
	var result profile.Row
	query := `SELECT * FROM user_profile WHERE id = $1`
	err := r.db.GetContext(ctx, &result, query, profileID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, common.ErrProfileNotFound
		}

		return nil, err
	}
	return &result, nil
}

func (r *Repository) Create(ctx context.Context, input profile.Row) (string, error) {
	query := `INSERT INTO user_profile (id, first_name, last_name, email, created_at, updated_at) VALUES (:id, :first_name, :last_name, :email, :created_at, :updated_at)`
	_, err := r.db.NamedExecContext(ctx, query, input)
	if err != nil {
		r.log.Errorw(ctx, "failed to create document", "err", err)
		return "", err
	}

	return input.ID, err
}
