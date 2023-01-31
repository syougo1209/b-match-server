package service

import (
	"context"

	"github.com/syougo1209/b-match-server/domain/model"
)

type AuthService interface {
	GenerateToken(ctx context.Context, uid model.User) ([]byte, error)
}
