package circle

import (
	"fmt"
	"log"
)

func Launch() {
	var area int

	log.Println("Укажите площадь круга:")
	circleArea, errCircleArea := fmt.Scanln(&area)
	if errCircleArea != nil {
		fmt.Printf("Вы не указали площадь круга")
		return
	}

	circleDiameter, circleLength := Diameter(circleArea), Length(circleArea)

	log.Println("Диаметр круга равен:", circleDiameter)
	log.Println("Длина круга равна:", circleLength)
}
