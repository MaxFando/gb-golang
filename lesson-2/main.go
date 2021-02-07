package main

import (
	"github.com/MaxFando/gb-golang/lesson-2/circle"
	"github.com/MaxFando/gb-golang/lesson-2/parse_inputs_int"
	"github.com/MaxFando/gb-golang/lesson-2/rectangle"
	"log"
	"os"
)

func main() {
	log.SetOutput(os.Stdout)

	rectangle.Launch()
	circle.Launch()
	parse_inputs_int.Launch()
}
