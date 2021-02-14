package main

import "github.com/MaxFando/gb-golang/lesson-8/resources"

func main() {
	configuration := resources.Configuration{}

	defer func() {
		configuration.CloseConnection()
	}()
	configuration.NewConnection()
}
