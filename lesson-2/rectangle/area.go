package rectangle

import (
	"errors"
	"fmt"
	"log"
)

// Area Calculate area of a rectangle
func Area(width, length float64) (float64, error) {
	if width == 0 {
		return 0, errors.New("Ширина = 0")
	}

	if length == 0 {
		return 0, errors.New("Длина = 0")
	}

	return width * length, nil
}

func Launch() {
	var width, length float64

	log.Println("Введите ширину:")
	if _, errWidth := fmt.Scanln(&width); errWidth != nil {
		log.Fatalln("Вы не указали ширину", errWidth)
	}

	log.Println("Введите длину:")
	if _, errLength := fmt.Scanln(&length); errLength != nil {
		log.Fatalln("Вы не указали длину", errLength)
	}

	if rectangleArea, errArea := Area(width, length); errArea != nil {
		log.Println(errArea)
	} else {
		log.Println("Площадь прямоугольника равна:", rectangleArea)
	}
}
