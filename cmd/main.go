package main

import (
	"kafka-connect-deployer/internal/config"
	"kafka-connect-deployer/internal/deployer"
	"log"
)

func main() {
	log.Print("Deployer starting....")

	c, err := config.New()
	if err != nil {
		log.Fatal("fail on start configuration", err)
	}

	d := deployer.New(c)
	d.Deploy()

}
