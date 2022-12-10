package models_test

import (
	"server-poc/pkg/models"
	"server-poc/pkg/testutils"
	"testing"
)

func TestMigrateAll(t *testing.T) {
	db := testutils.NewMockDB(t)
	err := models.MigrateAll(db)
	if err != nil {
		t.Fatal("failed to migrate tables:", err)
	}
}
