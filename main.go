package main

import (
	"flag"
	"log"
	"time"

	"server-poc/pkg/models"
	"server-poc/pkg/mqtt"
	"server-poc/pkg/testutils"
	"server-poc/services/datacollector"
	"server-poc/services/pumpcontroller"
	"server-poc/services/waterplanner"
	"server-poc/services/web"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

const (
	defaultBrokerAddr = "tcp://192.168.1.15:1883"
	mqttUsername      = "roslina"
	mqttPassword      = "smartcrops"
	defaultDBPath     = "artifacts/baza.db"
	defaultHTTPPort   = 8080
	restartTime       = time.Second * 2
)

func run() error {
	flagBrokerAddr := flag.String("broker-addr", defaultBrokerAddr, "Address of the MQTT broker.")
	flagWithBroker := flag.Bool("with-broker", false, "Starts an MQTT broker and connects all services to it.")
	flagDBPath := flag.String("db-path", defaultDBPath, "Path to the sqlite file")
	flagHTTPPort := flag.Int("http-port", defaultHTTPPort, "Sets the port on which the http server runs.")
	flag.Parse()

	/* -------------------------------- Database -------------------------------- */
	log.Println("Initializing database...")
	db, err := gorm.Open(sqlite.Open(*flagDBPath), &gorm.Config{})
	if err != nil {
		return err
	}

	log.Println("Performing database migrations...")
	err = models.MigrateAll(db)
	if err != nil {
		return err
	}

	/* ------------------------------- MQTT Broker ------------------------------ */
	if *flagWithBroker {
		broker, brokerAddr, err := testutils.StartMQTTBroker()
		if err != nil {
			return err
		}
		defer broker.Close()
		*flagBrokerAddr = brokerAddr
	}

	/* ------------------------------- MQTT Client ------------------------------ */
	log.Println("Initializing mqtt...")
	mqttClient, err := mqtt.Connect(*flagBrokerAddr, mqttUsername, mqttPassword)
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
	log.Println("Running http service on port", *flagHTTPPort)
	if err = web.Run(db, *flagHTTPPort); err != nil {
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
