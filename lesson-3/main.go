package main

import (
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	log.SetOutput(os.Stdout)

	fmt.Print("Enter first number: ")
	var firstNumber float64
	_, err := fmt.Scanln(&firstNumber)
	if err != nil {
		panic(fmt.Sprint("Error occurred", err.Error()))
	}

	fmt.Print("Enter second number: ")
	var secondNumber float64
	_, errSec := fmt.Scanln(&secondNumber)
	if errSec != nil {
		panic(fmt.Sprint("Error occurred", errSec.Error()))
	}

	fmt.Print("Enter operation (+, -, *, /, ^): ")
	var operation string
	_, errOp := fmt.Scanln(&operation)
	if errOp != nil {
		panic(fmt.Sprint("Error occurred", errOp.Error()))
	}

	var result float64
	switch operation {
	case "+":
		result = firstNumber + secondNumber
	case "-":
		result = firstNumber - secondNumber
	case "*":
		result = firstNumber * secondNumber
	case "/":
		result = firstNumber / secondNumber
	case "^":
		result = math.Pow(firstNumber, secondNumber)
	default:
		panic(fmt.Sprint("Operation is not defined"))
	}

	fmt.Printf("Operation result: %f\n", result)
}
