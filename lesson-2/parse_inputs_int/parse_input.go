package parse_inputs_int

import (
	"errors"
	"fmt"
)

func Parse(number int) (string, error) {
	parsedNumber := fmt.Sprintf("%v", number)

	if len(parsedNumber) > 3 || len(parsedNumber) < 3 {
		return "", errors.New("Укажите трехзначное число")
	}

	message := fmt.Sprintf("В числе %d - %s сотен, %s десятков, %s единиц", number, string(parsedNumber[0]), string(parsedNumber[1]), string(parsedNumber[2]))

	return message, nil
}
