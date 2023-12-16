package token

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Service represent token service
type Service struct {
}

// NewService create a new token service
func NewService() *Service {
	return &Service{}
}

var loginExpirationDuration = time.Duration(1) * time.Hour * 24
var jwtSigningMethod = jwt.SigningMethodHS256
var jwtSignatureKey = []byte("the secret of kalimdor")

// CreateToken create token
func (as *Service) CreateToken(ctx context.Context, req *CreateTokenRequest) (*Token, error) {
	claims := &Claims{
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: time.Now().Add(loginExpirationDuration).Unix(),
		},
		UserID: req.UserID,
		RoleID: req.RoleID,
	}

	jtoken := jwt.NewWithClaims(
		jwtSigningMethod,
		claims,
	)

	ts, err := jtoken.SignedString(jwtSignatureKey)
	if err != nil {
		return nil, err
	}
	token := &Token{
		AccessToken: ts,
	}
	return token, nil
}
