package pumpcontroller

import (
	"encoding/json"
	"fmt"
	"server-poc/pkg/mqtt"

	"gorm.io/gorm"
)

const (
	baseTopic = "command/"
	qos       = 2
	retained  = false
)

type Service interface {
	Send(Message, uint) error
}

type Message struct {
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

func (s service) Send(msg Message, destTankId uint) error {
	msgJson, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	return s.mqttClient.Pub(fmt.Sprintf("%s/%d", baseTopic, destTankId), qos, retained, msgJson)
}
