package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Account
	status struct {
		db *sqlx.DB
	}
)

var _ repository.Status = (*status)(nil)

// Create accout repository
func NewStatus(db *sqlx.DB) *status {
	return &status{db: db}
}

func (s *status) Create(ctx context.Context, tx *sqlx.Tx, st *object.Status) error {
	_, err := s.db.Exec("insert into status ( account_id, content, created_at) values (?,?,?)", st.AccountID, st.Content, st.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to insert account: %w", err)
	}
	return nil
}

// なぜここではtx *sqlx.Txこれが必要ないのだろう
func (s *status) FindStatusByID(ctx context.Context, id int) (*object.Status, error) {
	entity := new(object.Status)
	err := s.db.QueryRowxContext(ctx, "select * from status where id = ?", id).StructScan(entity)
	log.Print(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to find account from db: %w", err)
	}
	return entity, nil
}
