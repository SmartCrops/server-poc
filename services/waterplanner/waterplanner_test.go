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

func TestGetWaterPlanningData(t *testing.T) {
	// Setup
	is := is.New(t)
	db := setupEnviroment(t)

	sensorData := models.SensorData{}
	db.Last(&sensorData)
	log.Println("Last SensorData:", sensorData)

	data, err := waterplanner.GetWaterPlanningData(db, sensorData.DataCollectorSerialNumber)
	is.NoErr(err)

	log.Println("Data:", data)

	is.Equal(data.Lat, 140.0)
	is.Equal(data.Lon, 120.0)
	is.Equal(data.OptimalHumidity, 50.0)
	is.Equal(data.PumpControllerSerialNumber, "pumpcontroller2")
	is.Equal(data.PumpGpio, uint8(0))
	is.Equal(data.SoilHumidityAvg, 60.0)
}
