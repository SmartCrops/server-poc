package main

import (
	"log"
	"time"

	"server-poc/pkg/models"
	"server-poc/pkg/mqtt"
	"server-poc/services/datacollector"
	"server-poc/services/pumpcontroller"
	"server-poc/services/waterplanner"
	"server-poc/services/web"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

const (
	mqttURL      = "tcp://192.168.1.15:1883"
	mqttUsername = "roslina"
	mqttPassword = "smartcrops"
	dbPath       = "artifacts/baza.db"
	httpPort     = "8080"
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
	err = models.MigrateAll(db)
	if err != nil {
		return err
	}

	/* ------------------------------- MQTT Client ------------------------------ */
	log.Println("Initializing mqtt...")
	mqttClient, err := mqtt.Connect(mqttURL, mqttUsername, mqttPassword)
	if err != nil {
		return err
	}
	defer mqttClient.Close()

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

	/* ------------------------------- Web ------------------------------- */
	log.Println("Running http service on port", httpPort)
	if err = web.Run(db, httpPort); err != nil {
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
