package main

import (
	"server-poc/pkg/db"
	mobileapi "server-poc/pkg/mobile-api"
	"server-poc/pkg/mqtt"
	sensordata "server-poc/pkg/sensor-data"

	"log"
)

func main() {
	log.Println("Initializing database...")
	if err := db.Init(); err != nil {
		log.Fatal(err)
	}
	log.Println("Initializing mqtt...")
	if err := mqtt.Init(); err != nil {
		panic(err)
	}
	log.Println("Initializing sensordata service...")
	if err := sensordata.Init(); err != nil {
		panic(err)
	}
	log.Println("Running the mobile api...")
	if err := mobileapi.Run(); err != nil {
		panic(err)
	}
}
