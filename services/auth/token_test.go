package auth

import (
	"server-poc/pkg/models"
	"testing"
	"time"

	"github.com/matryer/is"
)

func TestTokens(t *testing.T) {
	const (
		username = "jankowalski"
		password = "some hash value"
		lifetime = time.Hour * 2
		secret   = "nocą tupta jeż"
	)
	is := is.New(t)

	user := models.User{
		Username:     username,
		PasswordHash: password,
	}

	tok, err := createToken(user, lifetime, []byte(secret))
	is.NoErr(err) // Should generate a token
	t.Log("generated a token:", tok)

	data, err := verifyToken(tok, []byte(secret))
	is.NoErr(err) // Should verify the token
	t.Log("decoded data:", data)
	is.Equal(data.Username, username) // Should get the correct user
}
