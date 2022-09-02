package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/syougo1209/b-match-server/domain/model"
)

type UserRepository struct {
	Db DbConnection
}

func (ur *UserRepository) FindByID(ctx context.Context, id model.UserID) (*model.User, error) {
	dto := &User{}
	query := `SELECT *
	    FROM user
			WHERE id = ?
	`

	if err := ur.Db.GetContext(ctx, dto, query, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("GetContext user by id=%d: %w", id, model.ErrNotFound)
		}
		return nil, fmt.Errorf("GetContext user by id=%d: %w", id, err)
	}
	user := &model.User{
		ID:   model.UserID(dto.ID),
		Name: dto.Name,
	}

	return user, nil
}

type User struct {
	ID   uint64 `db:"id"`
	Name string `db:"name"`
}
