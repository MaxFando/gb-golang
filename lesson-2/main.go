package main

import (
	"github.com/MaxFando/gb-golang/lesson-2/circle"
	"github.com/MaxFando/gb-golang/lesson-2/parse_inputs_int"
	"log"
	"os"

	"github.com/MaxFando/gb-golang/lesson-2/rectangle"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	rectangle.Launch()
	circle.Launch()
	parse_inputs_int.Launch()
}
