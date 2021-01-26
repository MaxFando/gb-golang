package main

import (
	"fmt"
	"log"
	"math"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	var a, b, res float64
	var op string

	fmt.Print("Введите первое число: ")
	_, err := fmt.Scanln(&a)
	if err != nil {
		log.Fatalln("Произошла ошибка", err.Error())
	}

	fmt.Print("Введите второе число: ")
	_, errSec := fmt.Scanln(&b)
	if errSec != nil {
		log.Fatalln("Произошла ошибка", errSec.Error())
	}

	fmt.Print("Введите арифметическую операцию (+, -, *, /, ^): ")
	_, errOp := fmt.Scanln(&op)
	if errOp != nil {
		log.Fatalln("Произошла ошибка", errOp.Error())
	}

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	case "^":
		res = math.Pow(a, b)
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}

	fmt.Printf("Результат выполнения операции: %f\n", res)
}
