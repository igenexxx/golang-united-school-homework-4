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

func getMatches(expression string) (error, []string) {
	var err error = nil

	if strings.Trim(expression, " ") == "" {
		return fmt.Errorf("%s", errorEmptyInput), nil
	}

	invalidExpr := regexp.MustCompile(`[^ \d+-]`)

	if invalidExpr.MatchString(expression) {
		return fmt.Errorf("invalid character"), nil
	}

	cleanExprRegex := regexp.MustCompile(`[^\d+-]`)
	matchesRegex := regexp.MustCompile(`(\+?-?\d+)([-+])(\+?-?\d+)`)

	cleanExpr := cleanExprRegex.ReplaceAllString(expression, "")
	matches := matchesRegex.FindStringSubmatch(cleanExpr)

	if matches == nil {
		return fmt.Errorf("no matches"), nil
	}

	if len(matches[1:]) != 3 {
		return fmt.Errorf("%s", errorNotTwoOperands), nil
	}

	if matches != nil {
		return nil, matches[1:]
	}

	return err, nil
}

func sum(a, b int) int {
	return a + b
}

func subtr(a, b int) int {
	return a - b
}

func parse(expression string) (string, error) {
	err, result := getMatches(expression)

	if err != nil {
		return "", err
	}

	operator := result[1]
	a, err := strconv.Atoi(result[0])
	b, err := strconv.Atoi(result[2])

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
