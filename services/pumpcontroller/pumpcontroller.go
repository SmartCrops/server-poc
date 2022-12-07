package pumpcontroller

import (
	"encoding/json"
	"fmt"
	"server-poc/pkg/mqtt"

	"gorm.io/gorm"
)

const (
	baseTopic = "command"
	qos       = 2
	retained  = false
)

type Service interface {
	Send(command PumpControllerCommand, destTankSerialNumber string) error
}

type PumpControllerCommand struct {
	PumpGpio  uint8  `json:"pumpGpio"`
	DurationS uint16 `json:"durationS"`
}

func Create(db *gorm.DB, mqttClient mqtt.Client) Service {
	s := service{
		db:         db,
		mqttClient: mqttClient,
	}

	return &s
}

type service struct {
	db         *gorm.DB
	mqttClient mqtt.Client
}

func (s service) Send(command PumpControllerCommand, destTankSerialNumber string) error {
	msgJson, err := json.Marshal(command)
	if err != nil {
		return err
	}

	return s.mqttClient.Pub(fmt.Sprintf("%s/%s", baseTopic, destTankSerialNumber), qos, retained, msgJson)
}
