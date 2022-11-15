package sensordata

import (
	"encoding/json"
	"log"

	db "server-poc/pkg/db"
	mqtt "server-poc/pkg/mqtt"
)

func handleNewData(msg mqtt.Msg) {
	type MessagePayload struct {
		Temp  float64 `json:"temp"`
		Pres  float64 `json:"pres"`
		Light float64 `json:"light"`
	}

	// Parse payload
	payload := MessagePayload{}
	err := json.Unmarshal(msg.Payload(), &payload)
	if err != nil {
		log.Default().Printf("failed to parse mqtt payload: %v", err)
		return
	}

	// Save data to the database
	data := SensorData{
		Temperature:   payload.Temp,
		Pressure:      payload.Pres,
		LightIntesity: payload.Light,
	}
	err = db.DB.Save(data).Error
	if err != nil {
		log.Default().Printf("failed to save data to the database: %v", err)
		return
	}
}
