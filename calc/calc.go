package calc

import (
	"errors"
)

//calculator is a structure to keep provided input values of 2 operands and math operation
type calculator struct{}

//NewCalculator is public method to create an instance of calculator structure
func NewCalculator() calculator {
	return calculator{}
}

//Calculate is the method to manage calculations, get 2 operands and math operation
func (c *calculator) Calculate(firstNumber int, secondNumber int, mathOperation string) (int, error) {
	switch mathOperation {
	case "+":
		return c.addition(firstNumber, secondNumber), nil
	case "-":
		return c.subtraction(firstNumber, secondNumber), nil
	case "*":
		return c.multiplication(firstNumber, secondNumber), nil
	case "/":
		res, err := c.division(firstNumber, secondNumber)
		return res, err
	default:
		return -1, errors.New("operation is not supported or is invalid, please try again")
	}
}

//addition method, gets in input operands to calculate
func (c *calculator) addition(firstNumber int, secondNumber int) int {
	return firstNumber + secondNumber
}

//subtraction method, gets in input operands to calculate
func (c *calculator) subtraction(firstNumber int, secondNumber int) int {
	return firstNumber - secondNumber
}

//multiplication method, gets in input operands to calculate
func (c *calculator) multiplication(firstNumber int, secondNumber int) int {
	return firstNumber * secondNumber
}

//division method, gets in input operands to calculate, if second operand is zero error message is printed
func (c *calculator) division(firstNumber int, secondNumber int) (int, error) {
	if secondNumber != 0 {
		return firstNumber / secondNumber, nil
	} else {
		return -1, errors.New("Division to zero is not possible, please try again with other second operand")
	}
}
