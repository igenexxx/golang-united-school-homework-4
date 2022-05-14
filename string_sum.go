package string_sum

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func getMatches(expression string) ([]string, error) {
	var err error = nil

	if strings.Trim(expression, " ") == "" {
		return nil, fmt.Errorf("%s", errorEmptyInput)
	}

	expressionWithoutSpaces := strings.ReplaceAll(expression, " ", "")

	matchesRegex := regexp.MustCompile(`^(-?.*?)([-+])(.*?)$`)

	matches := matchesRegex.FindStringSubmatch(expressionWithoutSpaces)

	if matches == nil || len(matches[1:]) != 3 {
		return nil, errorNotTwoOperands
	}

	if matches != nil {
		return matches[1:], nil
	}

	return nil, err
}

func sum(a, b int) int {
	return a + b
}

func subtr(a, b int) int {
	return a - b
}

func parse(expression string) (string, error) {
	result, err := getMatches(expression)

	if err != nil {
		return "", fmt.Errorf("%s", err)
	}

	operator := result[1]
	a, err := strconv.Atoi(result[0])
	if err != nil {
		return "", fmt.Errorf("%s", err)
	}

	b, err := strconv.Atoi(result[2])
	if err != nil {
		return "", fmt.Errorf("%s", err)
	}

	if operator == "+" {
		return strconv.Itoa(sum(a, b)), nil
	} else if operator == "-" {
		return strconv.Itoa(subtr(a, b)), nil
	}

	return "", nil
}

func StringSum(input string) (output string, err error) {
	output, err = parse(input)

	if err != nil {
		return "", err
	}

	return output, nil
}
