package auth

import (
	"fmt"
	"server-poc/pkg/models"
	"time"

	"github.com/golang-jwt/jwt"
)

type tokenClaims struct {
	jwt.StandardClaims
	User models.User `json:"user"`
}

func createToken(u models.User, lifetime time.Duration, secret []byte) (string, error) {
	t := jwt.New(jwt.GetSigningMethod("HS256"))
	t.Claims = &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(lifetime).Unix(),
		},
		u,
	}
	return t.SignedString(secret)
}

func verifyToken(tokStr string, secret []byte) (models.User, error) {
	tok, err := jwt.ParseWithClaims(tokStr, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return models.User{}, fmt.Errorf("failed to parse the token: %w", err)
	}
	claims, ok := tok.Claims.(*tokenClaims)
	if !ok {
		return models.User{}, fmt.Errorf("failed to cast token claims")
	}
	ok = claims.VerifyExpiresAt(time.Now().Unix(), true)
	if !ok {
		return models.User{}, fmt.Errorf("token expired")
	}
	return claims.User, nil
}
