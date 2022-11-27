package main

import (
	"log"
	"time"

	"server-poc/pkg/mobileapi"
	"server-poc/pkg/models"
	"server-poc/pkg/mqtt"
	"server-poc/pkg/services/datacollector"
	"server-poc/pkg/services/pumpcontroller"
	"server-poc/pkg/services/waterplanner"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

const (
	mqttURL      = "tcp://192.168.1.15:1883"
	mqttUsername = "roslina"
	mqttPassword = "smartcrops"
	dbPath       = "artifacts/baza.db"
	apiPort      = "8080"
	restartTime  = time.Second * 2
)

func run() error {
	/* -------------------------------- Database -------------------------------- */
	log.Println("Initializing database...")
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return err
	}

	log.Println("Performing database migrations...")
	err = db.AutoMigrate(models.SensorData{})
	if err != nil {
		return err
	}

	/* ------------------------------- MQTT Client ------------------------------ */
	log.Println("Initializing mqtt...")
	mqttClient, err := mqtt.Connect(mqttURL, mqttUsername, mqttPassword)
	if err != nil {
		return err
	}

	/* ----------------------------- Data Collector ----------------------------- */
	log.Println("Initializing datacollector service...")
	datacollectorService, err := datacollector.Start(mqttClient, db)
	if err != nil {
		return err
	}

	/* ----------------------------- Pump Controller ---------------------------- */
	log.Println("Initializing pumpcontroller service...")
	pumpcontrollerService := pumpcontroller.Create(db, mqttClient)

	/* ------------------------------ Water Planner ----------------------------- */
	log.Println("Initializing waterplanner service...")
	waterplanner.Start(db, mqttClient, datacollectorService, pumpcontrollerService)

	/* ------------------------------- Mobile Api ------------------------------- */
	log.Println("Running mobile api on port", apiPort)
	if err := mobileapi.Run(db, apiPort); err != nil {
		return err
	}

	return nil
}

func main() {
	for {
		err := run()
		log.Println("Run function exited:", err)
		log.Println("Scheduling next startup in", restartTime)
		time.Sleep(restartTime)
	}
}
