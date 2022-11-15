package sensordata

import (
	"fmt"
	db "server-poc/pkg/db"
	mqtt "server-poc/pkg/mqtt"
)

func Init() error {
	// Auto migrate database table
	err := db.DB.AutoMigrate(SensorData{})
	if err != nil {
		return fmt.Errorf("failed to migrate SensorData table: %w", err)
	}

	// Subscribe to a mqtt topic
	err = mqtt.Subscribe("smart-crops/sensor-data", 1, handleNewData)
	if err != nil {
		return fmt.Errorf("failed to subscribe to smart-crops/sensor-data: %w", err)
	}

	return nil
}
