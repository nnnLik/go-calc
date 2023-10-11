package main

import (
	"errors"
	"strconv"
	"strings"
)

func calculateExpression(input string) (interface{}, error) {
	tokens := strings.Fields(input)

	if len(tokens) != 3 {
		return nil, errors.New("Invalid expression format.")
	}

	operand1, operator, operand2 := tokens[0], tokens[1], tokens[2]

	if !isValidOperand(operand1) || !isValidOperand(operand2) {
		return nil, errors.New("Invalid input.")
	}

	result, err := performCalculation(operand1, operator, operand2)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func isValidOperand(operand string) bool {
	return (isRoman(operand) && isRoman(operand)) || (isArabic(operand) && isArabic(operand))
}

func performCalculation(operand1, operator, operand2 string) (interface{}, error) {
	var result interface{}

	if isArabic(operand1) {
		operand1Int, _ := strconv.Atoi(operand1)
		operand2Int, _ := strconv.Atoi(operand2)
		result, err := performArabicCalculation(operand1Int, operator, operand2Int)
		if err != nil {
			return nil, err
		}
		return result, nil
	}

	operand1Int, err := romanToArabic(operand1)
	if err != nil {
		return nil, err
	}

	operand2Int, err := romanToArabic(operand2)
	if err != nil {
		return nil, err
	}

	arabicResult, err := performArabicCalculation(operand1Int, operator, operand2Int)
	if err != nil {
		return nil, err
	}

	result, err = arabicToRoman(arabicResult)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func performArabicCalculation(operand1 int, operator string, operand2 int) (int, error) {
	switch operator {
	case "+":
		return operand1 + operand2, nil
	case "-":
		return operand1 - operand2, nil
	case "*":
		return operand1 * operand2, nil
	case "/":
		if operand2 == 0 {
			return 0, errors.New("Division by zero is not allowed.")
		}
		return operand1 / operand2, nil
	default:
		return 0, errors.New("Invalid operator. Please use +, -, *, or /.")
	}
}
