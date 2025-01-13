package document

import "context"

type Repository interface {
	FindByID(ctx context.Context, id string) (*Row, error)
	Create(ctx context.Context, payload Row) (string, error)
	Update(ctx context.Context, payload Row) (string, error)
}
