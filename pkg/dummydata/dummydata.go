package dummydata

import (
	"server-poc/pkg/models"

	"gorm.io/gorm"
)

func Populate(db *gorm.DB) {

	models.Installation{
		Name: "Kapusta",
		Lat:  52.237049,
		Lon:  21.017532,
	}.Save(db)

	models.Tank{

		InstallationID: 0,
	}.Save(db)

	models.Pump{
		TankID: 0,
		Gpio:   0,
	}.Save(db)
}
