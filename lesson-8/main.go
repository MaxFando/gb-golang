package main

import (
	"github.com/MaxFando/gb-golang/lesson-8/resources"
	"log"
	"os"
)

func main() {
	log.SetOutput(os.Stdin)

	configuration := resources.Configuration{}
	configuration.NewConnection()

	defer func() {
		configuration.CloseConnection()
	}()
}
