package database

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/syougo1209/b-match-server/application/usecase"
)

var txKey = struct{}{}

type tx struct {
	db *sqlx.DB
}

func NewTransaction(db *sqlx.DB) usecase.Transaction {
	return &tx{db: db}
}

func (t *tx) BeginTx(ctx context.Context, f func(ctx context.Context) (interface{}, error)) (interface{}, error) {
	tx, err := t.db.BeginTxx(ctx, nil)
	if err != nil {
		log.Print("failed to start transaction: %w", err)
		return nil, err
	}

	ctx = context.WithValue(ctx, &txKey, tx)
	v, err := f(ctx)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			log.Printf("failed to rollback: %v\n", rollbackErr)
		}
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			log.Printf("failed to rollback: %v\n", rollbackErr)
		}
		return nil, err
	}
	return v, nil
}

func GetTx(ctx context.Context) (*sqlx.Tx, bool) {
	tx, ok := ctx.Value(&txKey).(*sqlx.Tx)
	return tx, ok
}
