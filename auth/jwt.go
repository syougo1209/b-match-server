package auth

import (
	"context"
	_ "embed"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/syougo1209/b-match-server/domain/model"
)

//go:embed cert/secret.pem
var rawPrivateKey []byte

//go:embed cert/public.pem
var rawPubKey []byte

type JWTer struct {
	PrivateKey, PublicKey jwk.Key
	Store                 Store
}

type Store interface {
	Save(ctx context.Context, key string, uid model.UserID) error
	Load(ctx context.Context, key string) (model.UserID, error)
}

func NewJWTer(s Store) (*JWTer, error) {
	j := &JWTer{Store: s}
	privkey, err := jwk.ParseKey(rawPrivateKey, jwk.WithPEM(true))
	if err != nil {
		return nil, fmt.Errorf("failed in NewJWTer: private key: %w", err)
	}
	pubkey, err := jwk.ParseKey(rawPubKey, jwk.WithPEM(true))
	if err != nil {
		return nil, fmt.Errorf("failed in NewJWTer: public key: %w", err)
	}
	j.PrivateKey = privkey
	j.PublicKey = pubkey
	return j, nil
}

func (j *JWTer) GenerateToken(ctx context.Context, u model.User) ([]byte, error) {
	tok, err := jwt.NewBuilder().
		JwtID(uuid.New().String()).
		Issuer(`github.com/syougo1209/b-match-server`).
		Subject("access_token").
		IssuedAt(time.Now()).
		Build()

	if err != nil {
		return nil, fmt.Errorf("GetToken: failed to build token: %w", err)
	}
	if j.Store.Save(ctx, tok.JwtID(), u.ID); err != nil {
		return nil, err
	}

	signed, err := jwt.Sign(tok, jwt.WithKey(jwa.RS256, j.PrivateKey))
	if err != nil {
		return nil, err
	}

	return signed, nil
}

func (j *JWTer) GetToken(ctx context.Context, r *http.Request) (jwt.Token, error) {
	token, err := jwt.ParseRequest(r, jwt.WithKey(jwa.RS256, j.PublicKey))
	if err != nil {
		return nil, err
	}
	if _, err := j.Store.Load(ctx, token.JwtID()); err != nil {
		return nil, fmt.Errorf("GetToekn, %q : %w", token.JwtID(), err)
	}

	return token, nil
}

func (j *JWTer) FillContext(r *http.Request) (*http.Request, error) {
	token, err := j.GetToken(r.Context(), r)
	if err != nil {
		return nil, err
	}
	uid, err := j.Store.Load(r.Context(), token.JwtID())
	if err != nil {
		return nil, err
	}
	ctx := SetUserID(r.Context(), uid)

	clone := r.Clone(ctx)
	return clone, nil
}

type userIDKey struct{}

func SetUserID(ctx context.Context, uid model.UserID) context.Context {
	return context.WithValue(ctx, userIDKey{}, uid)
}
func GetUserID(ctx context.Context) (model.UserID, bool) {
	id, ok := ctx.Value(userIDKey{}).(model.UserID)
	return id, ok
}
