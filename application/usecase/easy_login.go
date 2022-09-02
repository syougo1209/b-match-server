package usecase

import (
	"context"
	"fmt"

	"github.com/syougo1209/b-match-server/domain/model"
	"github.com/syougo1209/b-match-server/domain/repository"
	"github.com/syougo1209/b-match-server/domain/service"
)

type EasyLogin interface {
	Call(ctx context.Context, uid model.UserID) (string, error)
}
type easyLogin struct {
	Repo   repository.UserRepository
	JwtGen service.AuthService
}

func NewEasyLogin(ur repository.UserRepository, jwt service.AuthService) *easyLogin {
	return &easyLogin{
		Repo:   ur,
		JwtGen: jwt,
	}
}
func (el *easyLogin) Call(ctx context.Context, uid model.UserID) (string, error) {
	u, err := el.Repo.FindByID(ctx, uid)
	if err != nil {
		return "", fmt.Errorf("easy login call findbyID %d: %w", uid, err)
	}
	jwt, err := el.JwtGen.GenerateToken(ctx, *u)
	if err != nil {
		return "", fmt.Errorf("easy login call failed to generate JWT: %w", err)
	}
	return string(jwt), nil
}
