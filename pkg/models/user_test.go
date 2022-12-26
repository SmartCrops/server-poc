package models_test

import (
	"server-poc/pkg/models"
	"server-poc/pkg/testutils"
	"testing"

	"github.com/matryer/is"
)

func TestUser(t *testing.T) {
	const (
		username = "jankowalski"
		password = "some hash value"
	)

	is := is.New(t)
	db := testutils.NewMockDB(t)

	data := models.User{
		Username:     username,
		PasswordHash: password,
	}

	t.Log("saving data:", data)
	is.NoErr(data.Save(db))

	var u models.User
	is.NoErr(u.FindByUsername(db, username))
	t.Log("user found by username:", u)
	is.Equal(u.Username, username)
	is.Equal(u.PasswordHash, password)

	is.NoErr(u.FindByCreds(db, username, password))
	t.Log("user found by creds:", u)
	is.Equal(u.Username, username)
	is.Equal(u.PasswordHash, password)
	id := u.ID
	u = models.User{}
	is.NoErr(u.GetByID(db, id))
	t.Log("user found by id:", u)
}
