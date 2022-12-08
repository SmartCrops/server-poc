package datacollector_test

import (
	"server-poc/pkg/models"
	"server-poc/pkg/mqtt"
	"server-poc/pkg/testutils"
	"server-poc/services/datacollector"
	"testing"
	"time"

	"github.com/matryer/is"
	"gorm.io/gorm"
)

func setupEnviroment(t *testing.T) (mqtt.Client, mqtt.Broker, *gorm.DB, datacollector.Service) {
	is := is.New(t)
	mqttClient, mqttBroker := mqtt.Mock(t)
	db := testutils.NewMockDB(t)
	service, err := datacollector.Start(mqttClient, db)
	is.NoErr(err) // DataCollector service should start
	return mqttClient, mqttBroker, db, service
}

func TestNotification(t *testing.T) {
	// Setup
	is := is.New(t)
	mqttClient, mqttBroker, db, service := setupEnviroment(t)
	defer mqttClient.Close()
	defer mqttBroker.Close()

	// Create a notification handler
	notification := make(chan models.SensorData, 1)
	service.ListenForNewData(func(sd models.SensorData) {
		notification <- sd
	})

	// Send new data to the mqtt broker
	payload := []byte(`{"soilHumidity": 66.0, "temperature":21.5, "lightIntensity": 123.0, "serialNumber": "xyz123"}`)
	err := mqttBroker.Publish("sensors/1", payload, false)
	is.NoErr(err) // Data should be published on the mqtt broker

	// Wait for a notification
	select {
	case <-time.After(time.Millisecond * 10):
		is.Fail() // Timeout occurred while waiting for a notification
	case <-notification:
	}

	// Check the database
	var data models.SensorData
	queryResult := db.First(&data)
	is.NoErr(queryResult.Error)            // Should query the database for sensordata
	is.True(queryResult.RowsAffected == 1) // Should find exactly one row
	// Data in the database should have correct fields
	is.True(data.SoilHumidity == 66.0)
	is.True(data.Temperature == 21.5)
	is.True(data.LightIntesity == 123.0)
	is.True(data.DataCollectorSerialNumber == "xyz123")
}

func TestInvalidData(t *testing.T) {
	// Setup
	is := is.New(t)
	mqttClient, mqttBroker, db, service := setupEnviroment(t)
	defer mqttClient.Close()
	defer mqttBroker.Close()

	// Create a notification handler
	notification := make(chan models.SensorData, 1)
	service.ListenForNewData(func(sd models.SensorData) {
		notification <- sd
	})

	// Send new data to the mqtt broker
	payload := []byte(`this is an invalid json`)
	err := mqttBroker.Publish("sensors/1", payload, false)
	is.NoErr(err) // Data should be published on the mqtt broker

	// Wait for a notification
	select {
	case <-time.After(time.Millisecond * 10):
	case <-notification:
		is.Fail() // Got a notification on invalid data
	}

	// Check the database
	var data []models.SensorData
	queryResult := db.Find(&data)
	is.NoErr(queryResult.Error)            // Should query database for sensordata
	is.True(queryResult.RowsAffected == 0) // Shouldn't find any rows
	is.True(len(data) == 0)                // Shouldn't return any results
}
