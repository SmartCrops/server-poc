package testutils

import (
	"server-poc/pkg/models"
	"testing"
	"time"

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

func PopulateMockData(t *testing.T, db *gorm.DB) {
	is := is.New(t)

	// Create User Frank
	user := models.User{
		Username:     "Frank",
		PasswordHash: "xyz123",
	}
	is.NoErr(user.Save(db))

	// Create 2 installations for user Frank
	installation1 := models.Installation{
		Name:            "Pole kapusty",
		Lat:             150.0,
		Lon:             120.0,
		OptimalHumidity: 60.0,
		UserID:          1,
	}
	is.NoErr(installation1.Save(db))

	installation2 := models.Installation{
		Name:            "Pole ziemniaków",
		Lat:             140.0,
		Lon:             120.0,
		OptimalHumidity: 50.0,
		UserID:          1,
	}
	is.NoErr(installation2.Save(db))

	// Create PumpController for every Installation
	pumpController1 := models.PumpController{
		SerialNumber:   "pumpcontroller1",
		InstallationID: 1,
	}
	is.NoErr(pumpController1.Save(db))
	pumpController2 := models.PumpController{
		SerialNumber:   "pumpcontroller2",
		InstallationID: 2,
	}
	is.NoErr(pumpController2.Save(db))

	// Create a Pump for every PumpController
	pump1 := models.Pump{
		SerialNumber:               "pump1",
		Gpio:                       0,
		PumpControllerSerialNumber: "pumpcontroller1",
	}
	is.NoErr(pump1.Save(db))

	pump2 := models.Pump{
		SerialNumber:               "pump2",
		Gpio:                       0,
		PumpControllerSerialNumber: "pumpcontroller2",
	}
	is.NoErr(pump2.Save(db))

	// Create 2 collectors for pump1
	collector11 := models.DataCollector{
		SerialNumber:     "datacollector11",
		PumpSerialNumber: "pump1",
	}
	is.NoErr(collector11.Save(db))
	collector12 := models.DataCollector{
		SerialNumber:     "datacollector12",
		PumpSerialNumber: "pump1",
	}
	is.NoErr(collector12.Save(db))

	// Create 2 collectors for pump2
	collector21 := models.DataCollector{
		SerialNumber:     "datacollector21",
		PumpSerialNumber: "pump2",
	}
	is.NoErr(collector21.Save(db))
	collector22 := models.DataCollector{
		SerialNumber:     "datacollector22",
		PumpSerialNumber: "pump2",
	}
	is.NoErr(collector22.Save(db))

	// Create SensorData (2 controllers x 1 pump x 2 collectors x 2 records == 8 records)
	data111 := models.SensorData{
		SoilHumidity:              80.0,
		Temperature:               21.5,
		LightIntesity:             123.0,
		DataCollectorSerialNumber: "datacollector11",
	}
	is.NoErr(data111.Save(db))
	data112 := models.SensorData{
		SoilHumidity:              70.0,
		Temperature:               21.5,
		LightIntesity:             123.0,
		DataCollectorSerialNumber: "datacollector11",
	}
	is.NoErr(data112.Save(db))

	data121 := models.SensorData{
		SoilHumidity:              70.0,
		Temperature:               21.5,
		LightIntesity:             123.0,
		DataCollectorSerialNumber: "datacollector12",
	}
	is.NoErr(data121.Save(db))
	data122 := models.SensorData{
		SoilHumidity:              60.0,
		Temperature:               21.5,
		LightIntesity:             123.0,
		DataCollectorSerialNumber: "datacollector12",
	}
	is.NoErr(data122.Save(db))

	data211 := models.SensorData{
		SoilHumidity:              80.0,
		Temperature:               21.5,
		LightIntesity:             123.0,
		DataCollectorSerialNumber: "datacollector21",
	}
	is.NoErr(data211.Save(db))

	data221 := models.SensorData{
		SoilHumidity:              80.0,
		Temperature:               21.5,
		LightIntesity:             123.0,
		DataCollectorSerialNumber: "datacollector22",
	}
	is.NoErr(data221.Save(db))

	time.Sleep(time.Millisecond) // Make sure following records are inserted last for testing
	data212 := models.SensorData{
		SoilHumidity:              50.0,
		Temperature:               21.5,
		LightIntesity:             123.0,
		DataCollectorSerialNumber: "datacollector21",
	}
	is.NoErr(data212.Save(db))
	data222 := models.SensorData{
		SoilHumidity:              70.0,
		Temperature:               21.5,
		LightIntesity:             123.0,
		DataCollectorSerialNumber: "datacollector22",
	}
	is.NoErr(data222.Save(db))
}
