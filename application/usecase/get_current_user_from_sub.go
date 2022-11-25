package usecase

import (
	"context"
	"fmt"

	"github.com/syougo1209/b-match-server/domain/model"
	"github.com/syougo1209/b-match-server/domain/repository"
)

type GetCurrentUserFromSub interface {
	Call(ctx context.Context, sub string) (model.UserID, error)
}

type getCurrentUserFromSub struct {
	userRepo repository.UserRepository
}

func NewGetCurrentUserFromSub(
	ur repository.UserRepository,
) *getCurrentUserFromSub {
	return &getCurrentUserFromSub{
		userRepo: ur,
	}
}

func (a *getCurrentUserFromSub) Call(ctx context.Context, sub string) (*model.UserID, error) {
	uid, err := a.userRepo.FindBySub(ctx, sub)
	if err != nil {
		return nil, fmt.Errorf("failed to find user from sub: %w", err)
	}
	return uid, nil
}
