package auth_test

import (
	"server-poc/pkg/testutils"
	"server-poc/services/auth"
	"testing"
	"time"

	"github.com/matryer/is"
)

// func TestCreateToken(t *testing.T) {
// 	tok, err := createToken()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	fmt.Println(tok)
// }
// func TestParseToken(t *testing.T) {
// 	tok, err := createToken()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	err = parseToken(tok)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }

func TestAuthFlow(t *testing.T) {
	const (
		secret   = "nocą tupta jeż"
		lifetime = time.Hour * 2
		username = "jankowalski"
		password = "janek123"
	)

	is := is.New(t)
	db := testutils.NewMockDB(t)
	authService := auth.Service{
		DB:            db,
		Secret:        []byte(secret),
		TokenLifetime: lifetime,
	}

	_, err := authService.Authenticate("invalid token")
	is.True(err != nil) // Authenticating an invalid token should fail

	_, err = authService.Login("wrong", "wrong")
	is.True(err != nil) // Trying to login with wrong credentials should fail

	is.NoErr(authService.Register(username, password)) // Should register a new user

	err = authService.Register(username, "doesnt matter")
	is.True(err != nil) // Should fail if username already exists

	tok, err := authService.Login(username, password)
	is.NoErr(err) // Should login correctly

	user, err := authService.Authenticate(tok)
	t.Log("user data: ", user)
	is.NoErr(err)                     // Should accept a valid token
	is.Equal(user.Username, username) // Should get the right user

	_, err = authService.Login(username, "wrong")
	is.True(err != nil) // Should fail if password is wrong

	_, err = authService.Login("wrong", password)
	is.True(err != nil) // Should fail if username is wrong
}
