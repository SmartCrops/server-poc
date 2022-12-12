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
	db := testutils.NewMockDB(t)
	testutils.PopulateMockData(t, db)
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
