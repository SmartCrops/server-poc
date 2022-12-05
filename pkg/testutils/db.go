package testutils

import (
	"server-poc/pkg/models"
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/matryer/is"
	"gorm.io/gorm"
)

func NewMockDB(t *testing.T) *gorm.DB {
	is := is.New(t)
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	is.NoErr(err) // Database should open
	err = models.MigrateAll(db)
	is.NoErr(err) // Database should automigrate
	return db
}
