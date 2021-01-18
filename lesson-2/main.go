package main

import (
	"fmt"
	"github.com/MaxFando/gb-golang/lesson-2/circle"
	"github.com/MaxFando/gb-golang/lesson-2/parse_inputs_int"

	"github.com/MaxFando/gb-golang/lesson-2/rectangle"
)

func main() {
	var width, length float64
	var number int

	fmt.Print("Введите ширину:")
	_, errWidth := fmt.Scanln(&width)

	if errWidth != nil {
		fmt.Printf("Вы не указали ширину")
	}

	fmt.Print("Введите длину:")
	_, errLength := fmt.Scanln(&length)

	if errLength != nil {
		fmt.Printf("Вы не указали длину")
	}

	rectangleArea, errArea := rectangle.Area(width, length)
	if errArea != nil {
		fmt.Printf(errArea.Error())
	} else {
		fmt.Println("Площадь прямоугольника равна:", rectangleArea)
	}

	fmt.Print("Укажите площадь круга:")
	circleArea, errCircleArea := fmt.Scanln(&width)

	if errCircleArea != nil {
		fmt.Printf("Вы не указали площадь круга")
		return
	}

	circleDiameter := circle.Diameter(circleArea)
	circleLength := circle.Length(circleArea)

	fmt.Println("Диаметр круга равен:", circleDiameter)
	fmt.Println("Длина круга равна:", circleLength)

	fmt.Print("Введите трехзначное число:")
	_, errNumber := fmt.Scanln(&number)

	if errNumber != nil {
		fmt.Printf("Не хочешь не надо")
		return
	}

	numMessage, errNumMessage := parse_inputs_int.Parse(number)
	if errNumMessage != nil {
		fmt.Printf(errNumMessage.Error())
	} else {
		fmt.Println(numMessage)
	}
}
