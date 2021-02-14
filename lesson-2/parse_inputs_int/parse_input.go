package parse_inputs_int

import (
	"errors"
	"fmt"
	"log"
)

//Parse Выводит кол-во сотен, десятков и единиц
func Parse(number uint32) (string, error) {
	parsedNumber := fmt.Sprintf("%v", number)

	if len(parsedNumber) > 3 || len(parsedNumber) < 3 {
		return "", errors.New("Укажите трехзначное число")
	}

	message := fmt.Sprintf("В числе %d - %s сотен, %s десятков, %s единиц",
		number, string(parsedNumber[0]),
		string(parsedNumber[1]), string(parsedNumber[2]))

	return message, nil
}

func Launch() {
	var number uint32

	log.Print("Введите трехзначное число:")
	if _, errNumber := fmt.Scanln(&number); errNumber != nil {
		fmt.Println("Не хочешь не надо", errNumber.Error())
		return
	}

	if numMessage, errNumMessage := Parse(number); errNumMessage != nil {
		fmt.Printf(errNumMessage.Error())
	} else {
		fmt.Println(numMessage)
	}
}
