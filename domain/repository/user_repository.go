package repository

import (
	"context"

	"github.com/syougo1209/b-match-server/domain/model"
)

type UserRepository interface {
	FindByID(ctx context.Context, uid model.UserID) (*model.User, error)
	FindBySub(ctx context.Context, sub string) (*model.UserID, error)
}
