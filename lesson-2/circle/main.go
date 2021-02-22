package circle

import (
	"fmt"
	"log"
	"math"
)

//Diameter
func Diameter(area int) float64 {
	return 2 * math.Sqrt(float64(area)/math.Pi)
}

//Length
func Length(area int) float64 {
	return 2 * math.Pi * Diameter(area)
}

func Launch() {
	var area uint32

	log.Println("Укажите площадь круга:")
	circleArea, errCircleArea := fmt.Scanln(&area)
	if errCircleArea != nil {
		fmt.Println("Вы не указали площадь круга", errCircleArea.Error())
		return
	}

	circleDiameter, circleLength := Diameter(circleArea), Length(circleArea)

	log.Println("Диаметр круга равен:", circleDiameter)
	log.Println("Длина круга равна:", circleLength)
}
