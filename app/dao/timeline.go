package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	timeline struct {
		db *sqlx.DB
	}
)

var _ repository.Timeline = (*timeline)(nil)

func NewTimeline(db *sqlx.DB) *timeline {
	return &timeline{db: db}
}

func (t *timeline) TimeFindStatusByIDsfalse(ctx context.Context, id int, limit int) ([]*object.Status, error) {
	var entities []*object.Status
	rows, err := t.db.QueryContext(ctx, "SELECT * FROM status WHERE id BETWEEN ? AND ?", id, id+limit)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to find status from db: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		entity := new(object.Status)
		if err := rows.Scan(
			&entity.ID,
			&entity.AccountID,
			&entity.Content,
			&entity.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan status row: %w", err)
		}
		entities = append(entities, entity)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}
	return entities, nil
}
