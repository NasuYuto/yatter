package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Timeline interface {
	TimeFindStatusByIDsfalse(ctx context.Context, id int, limit int) ([]*object.Status, error)
}
