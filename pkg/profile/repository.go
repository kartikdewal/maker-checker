package profile

import "context"

type Repository interface {
	FindByID(ctx context.Context, id string) (*Row, error)
	Create(ctx context.Context, input Row) (string, error)
}
