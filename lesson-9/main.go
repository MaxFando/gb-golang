package main

import (
	decoder "github.com/MaxFando/gb-golang/lesson-9/resources"
)

func main() {
	decoder.JsonDecode("app.json")
	decoder.YmlDecode("app.yml")
}
