package usecase

import (
	"context"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type Status interface {
	Create(ctx context.Context, content string, account_id int) (*CreateStatusDTO, error)
	FindStatusByID(ctx context.Context, id int) (*GetStatusDTO, error)
}

type status struct {
	db         *sqlx.DB
	statusRepo repository.Status
}

type CreateStatusDTO struct {
	Status *object.Status
}

type GetStatusDTO struct {
	Status *object.Status
}

var _ Status = (*status)(nil)

func NewStatus(db *sqlx.DB, statusRepo repository.Status) *status {
	return &status{
		db:         db,
		statusRepo: statusRepo,
	}
}

func (s *status) Create(ctx context.Context, content string, account_id int) (*CreateStatusDTO, error) {
	//err書かなくていいの？後で確認
	st := object.NewStatus(content, account_id)
	tx, err := s.db.Beginx()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}

		tx.Commit()
	}()

	if err := s.statusRepo.Create(ctx, tx, st); err != nil {
		return nil, err
	}

	return &CreateStatusDTO{
		Status: st,
	}, nil
}

func (s *status) FindStatusByID(ctx context.Context, id int) (*GetStatusDTO, error) {
	tx, err := s.db.Beginx()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}

		tx.Commit()
	}()
	dto, err := s.statusRepo.FindStatusByID(ctx, id)
	if err != nil {

		return nil, err
	}

	return &GetStatusDTO{
		Status: dto,
	}, nil
}
