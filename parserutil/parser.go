package parserutil

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

func ParseString(inputString string) (int, int, string, error) {
	operationRegex := regexp.MustCompile(`([0-9]*)([+\-*/])([0-9]*)=.*`)

	submatch := operationRegex.FindAllStringSubmatch(inputString, -1)
	if len(submatch) == 0 {
		return 0, 0, "", errors.New("string doesn't match pattern: a+b")
	}

	firstNumber, err := strconv.Atoi(submatch[0][1])
	if err != nil {
		return 0, 0, "", err
	}
	secondNumber, err := strconv.Atoi(submatch[0][3])
	if err != nil {
		return 0, 0, "", err
	}
	operand := submatch[0][2]

	fmt.Println(firstNumber)
	fmt.Println(secondNumber)

	return firstNumber, secondNumber, operand, nil
}
