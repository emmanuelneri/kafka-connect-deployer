package main

import (
	"kafka-connect-deployer/internal/config"
	"kafka-connect-deployer/internal/deployer"
	"log"
	"time"
)

func main() {
	log.Print("Deployer starting....")

	c, err := config.New()
	if err != nil {
		log.Panic(err)
	}

	if c.WaitStartTime > time.Nanosecond {
		log.Printf("waiting %s before start", c.WaitStartTime)
		time.Sleep(c.WaitStartTime)
	}

	d := deployer.New(c)
	d.Deploy()

}
