package waterplanner_test

import (
	"log"
	"server-poc/pkg/models"
	"server-poc/pkg/testutils"
	"server-poc/services/waterplanner"
	"testing"

	"github.com/matryer/is"
	"gorm.io/gorm"
)

func setupEnviroment(t *testing.T) *gorm.DB {
	// Create MockDB
	is := is.New(t)
	db := testutils.NewMockDB(t)

	// Populate with MockData
	user := models.User{
		Username:     "Frank",
		PasswordHash: "xyz123",
	}
	is.NoErr(user.Save(db))
	installation := models.Installation{
		Name:            "Pole kapusty",
		Lat:             150.0,
		Lon:             120.0,
		OptimalHumidity: 60.0,
		UserID:          1,
	}
	is.NoErr(installation.Save(db))
	pumpController := models.PumpController{
		SerialNumber:   "pumpcontroller1",
		InstallationID: 1,
	}
	is.NoErr(pumpController.Save(db))
	pump := models.Pump{
		SerialNumber:     "pump1",
		Gpio:             0,
		PumpControllerID: 1,
	}
	is.NoErr(pump.Save(db))
	collector1 := models.DataCollector{
		PumpID:       1,
		SerialNumber: "datacollector1",
	}
	is.NoErr(collector1.Save(db))
	collector2 := models.DataCollector{
		PumpID:       1,
		SerialNumber: "datacollector2",
	}
	is.NoErr(collector2.Save(db))
	data1 := models.SensorData{
		SoilHumidity:              50.0,
		Temperature:               21.5,
		LightIntesity:             123.0,
		DataCollectorSerialNumber: "datacollector1",
	}
	is.NoErr(data1.Save(db))
	data2 := models.SensorData{
		SoilHumidity:              70.0,
		Temperature:               21.5,
		LightIntesity:             123.0,
		DataCollectorSerialNumber: "datacollector2",
	}
	is.NoErr(data2.Save(db))
	return db
}

func TestWaterPlanningQuery(t *testing.T) {
	// Setup
	is := is.New(t)
	db := setupEnviroment(t)

	oldData, err := waterplanner.GetWaterPlanningData("datacollector1", db)
	log.Println(oldData, "brooo")
	is.NoErr(err)
	is.Equal(oldData.SoilHumidityAvg, oldData.OptimalHumidity)

	newSensorData := models.SensorData{
		SoilHumidity:              30.0,
		Temperature:               21.5,
		LightIntesity:             123.0,
		DataCollectorSerialNumber: "datacollector2",
	}

	newData, err := waterplanner.GetWaterPlanningData(newSensorData.DataCollectorSerialNumber, db)
	is.NoErr(err)
	is.Equal(newData.Lat, 150.0)
	is.Equal(newData.Lon, 120.0)
	is.Equal(newData.OptimalHumidity, 60.0)
	is.Equal(newData.PumpControllerSerialNumber, "pumpcontroller1")
	is.Equal(newData.PumpGpio, 0)
	is.Equal(newData.SoilHumidityAvg, (50.0+30.0)/2) // newSensorData updates avgHum from (50.0 + 70.0)/2 to (50.0+30.0)/2
}
