package datacollector

import (
	"encoding/json"
	"log"

	"server-poc/pkg/mqtt"
	"server-poc/pkg/sensordata"

	"gorm.io/gorm"
)

func Start(mqttClient mqtt.Client, db *gorm.DB) error {
	err := mqttClient.Sub("sensors/*", 1, func(msg mqtt.Msg) {
		handleData(db, msg)
	})
	return err
}

func handleData(db *gorm.DB, msg mqtt.Msg) {
	type MessagePayload struct {
		Temp     float64 `json:"temp"`
		Pres     float64 `json:"pres"`
		Light    float64 `json:"light"`
		SensorID int     `json:"sensorId"`
	}
	var payload MessagePayload
	if err := json.Unmarshal(msg.Payload(), &payload); err != nil {
		log.Println("Received malformed sensordata:", err)
		return
	}
	data := sensordata.SensorData{
		Temperature:   payload.Temp,
		Pressure:      payload.Pres,
		LightIntesity: payload.Light,
		SensorID:      payload.SensorID,
	}
	if err := data.Save(db); err != nil {
		log.Println("Failed to save sensor data to the database:", err)
	}
}
