package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"

	"github.com/jmoiron/sqlx"
)

type Status interface {
	FindStatusByID(ctx context.Context, id int) (*object.Status, error)
	Create(ctx context.Context, tx *sqlx.Tx, st *object.Status) error
}
