package datacollector_test

import (
	"server-poc/pkg/datacollector"
	"server-poc/pkg/mqtt"
	"server-poc/pkg/sensordata"
	"testing"
	"time"

	"github.com/glebarez/sqlite"
	"github.com/matryer/is"
	"gorm.io/gorm"
)

func setupEnviroment(t *testing.T) (mqtt.Client, mqtt.Broker, *gorm.DB, datacollector.Service) {
	is := is.New(t)
	mqttClient, mqttBroker := mqtt.Mock(t)
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	is.NoErr(err) // Database should open
	err = db.AutoMigrate(&sensordata.SensorData{})
	is.NoErr(err) // Database should automigrate
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
	notification := make(chan sensordata.SensorData, 1)
	service.ListenForNewData(func(sd sensordata.SensorData) {
		notification <- sd
	})

	// Send new data to the mqtt broker
	payload := []byte(`{"temp":21.5, "pres":1200.1, "light":123, "sensorId":1}`)
	err := mqttBroker.Publish("sensors/1", payload, false)
	is.NoErr(err) // Data should be published on the mqtt broker

	// Wait for a notification
	select {
	case <-time.After(time.Millisecond * 10):
		is.Fail() // Timeout occurred while waiting for a notification
	case <-notification:
	}

	// Check the database
	var data sensordata.SensorData
	queryResult := db.First(&data)
	is.NoErr(queryResult.Error)            // Should query the database for sensordata
	is.True(queryResult.RowsAffected == 1) // Should find exactly one row
	is.True(data.SensorID == 1)            // Data in the database should have correct fields
}

func TestInvalidData(t *testing.T) {
	// Setup
	is := is.New(t)
	mqttClient, mqttBroker, db, service := setupEnviroment(t)
	defer mqttClient.Close()
	defer mqttBroker.Close()

	// Create a notification handler
	notification := make(chan sensordata.SensorData, 1)
	service.ListenForNewData(func(sd sensordata.SensorData) {
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
	var data []sensordata.SensorData
	queryResult := db.Find(&data)
	is.NoErr(queryResult.Error)            // Should query database for sensordata
	is.True(queryResult.RowsAffected == 0) // Shouldn't find any rows
	is.True(len(data) == 0)                // Shouldn't return any results
}
