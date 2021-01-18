package main

import (
	"fmt"
	"github.com/MaxFando/gb-golang/lesson-2/circle"

	"github.com/MaxFando/gb-golang/lesson-2/rectangle"
)

func main() {
	var width, length float64

	fmt.Print("Enter width:")
	_, errWidth := fmt.Scanln(&width)

	if errWidth != nil {
		fmt.Printf("You didn't enter the width")
	}

	fmt.Print("Enter length:")
	_, errLength := fmt.Scanln(&length)

	if errLength != nil {
		fmt.Printf("You didn't enter the length")
	}

	rectangleArea, errArea := rectangle.Area(width, length)
	if errArea != nil {
		fmt.Printf(errArea.Error())
	} else {
		fmt.Println("Area of a rectangle is:", rectangleArea)
	}

	fmt.Print("Enter circle area:")
	circleArea, errCircleArea := fmt.Scanln(&width)

	if errCircleArea != nil {
		fmt.Printf("You didn't enter the width")
		return
	}

	circleDiameter := circle.Diameter(circleArea)
	circleLength := circle.Length(circleArea)

	fmt.Println("Circle diameter is:", circleDiameter)
	fmt.Println("Circle length is:", circleLength)
}
