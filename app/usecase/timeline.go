package usecase

import (
	"context"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type Timeline interface {
	Get(ctx context.Context, id int, limit int, flag bool) (*GetTimeDTO, error)
}

type timeline struct {
	db       *sqlx.DB
	timeRepo repository.Timeline
}

type GetTimeDTO struct {
	Timeline []*object.Status
}

var _ Timeline = (*timeline)(nil)

func NewTimeline(db *sqlx.DB, timeRepo repository.Timeline) *timeline {
	return &timeline{
		db:       db,
		timeRepo: timeRepo,
	}
}

func (t *timeline) Get(ctx context.Context, id int, limit int, flag bool) (*GetTimeDTO, error) {
	tx, err := t.db.Beginx()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}

		tx.Commit()
	}()
	if flag == false {
		dto, err := t.timeRepo.TimeFindStatusByIDsfalse(ctx, id, limit)
		if err != nil {
			return nil, err
		}
		return &GetTimeDTO{
			Timeline: dto,
		}, nil
	} else { //後ほど記述
		return nil, nil
	}
}
