package main

import (
	"fmt"
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
	var width, length float64
	var number int

	log.Println("Введите ширину:")
	if _, errWidth := fmt.Scanln(&width); errWidth != nil {
		log.Fatalln("Вы не указали ширину", errWidth)
	}

	log.Println("Введите длину:")
	if _, errLength := fmt.Scanln(&length); errLength != nil {
		log.Fatalln("Вы не указали длину", errLength)
	}

	if rectangleArea, errArea := rectangle.Area(width, length); errArea != nil {
		log.Println(errArea)
	} else {
		log.Println("Площадь прямоугольника равна:", rectangleArea)
	}

	log.Println("Укажите площадь круга:")
	circleArea, errCircleArea := fmt.Scanln(&width)
	if errCircleArea != nil {
		fmt.Printf("Вы не указали площадь круга")
		return
	}

	circleDiameter, circleLength := circle.Diameter(circleArea), circle.Length(circleArea)

	log.Println("Диаметр круга равен:", circleDiameter)
	log.Println("Длина круга равна:", circleLength)

	log.Print("Введите трехзначное число:")
	if _, errNumber := fmt.Scanln(&number); errNumber != nil {
		fmt.Printf("Не хочешь не надо")
		return
	}

	if numMessage, errNumMessage := parse_inputs_int.Parse(number); errNumMessage != nil {
		fmt.Printf(errNumMessage.Error())
	} else {
		fmt.Println(numMessage)
	}
}
