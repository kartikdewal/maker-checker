package document

import (
	"time"
)

type Row struct {
	ID          string    `db:"id"`
	Description string    `db:"description"`
	CreatorID   string    `db:"creator_id"`
	Status      string    `db:"status"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
