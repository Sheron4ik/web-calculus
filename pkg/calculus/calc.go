package calculus

import (
	"strconv"

	"github.com/Sheron4ik/web-calculus/pkg/errors"
)

var priority = map[string]int{
	"(": 1,
	"+": 2,
	"-": 2,
	"*": 3,
	"/": 3,
}

func isOperation(op string) bool {
	return op == "+" || op == "-" || op == "*" || op == "/"
}

func applyOperation(first, second float64, op string) (float64, error) {
	switch op {
	case "+":
		return first + second, nil
	case "-":
		return first - second, nil
	case "*":
		return first * second, nil
	case "/":
		if second == 0 {
			return 0, errors.ErrDivideByZero
		}
		return first / second, nil
	}
	return 0, errors.ErrUnknownOperations
}

func parse(expression string) ([]string, error) {
	polka := make([]string, 0)
	sign_stack := make([]string, 0)
	for _, ch := range expression {
		switch {
		case isOperation(string(ch)):
			for len(sign_stack) > 0 && priority[sign_stack[len(sign_stack)-1]] >= priority[string(ch)] {
				polka = append(polka, sign_stack[len(sign_stack)-1])
				sign_stack = sign_stack[:len(sign_stack)-1]
			}
			sign_stack = append(sign_stack, string(ch))
		case ch == rune('('):
			sign_stack = append(sign_stack, string(ch))
		case ch == rune(')'):
			for len(sign_stack) > 0 && sign_stack[len(sign_stack)-1] != "(" {
				polka = append(polka, sign_stack[len(sign_stack)-1])
				sign_stack = sign_stack[:len(sign_stack)-1]
			}
			if len(sign_stack) > 0 && sign_stack[len(sign_stack)-1] == "(" {
				sign_stack = sign_stack[:len(sign_stack)-1]
			} else {
				return nil, errors.ErrMissingOpenBracket
			}
		default:
			polka = append(polka, string(ch))
		}
	}
	for len(sign_stack) > 0 {
		if sign_stack[len(sign_stack)-1] == "(" {
			return nil, errors.ErrMissingCloseBracket
		}
		polka = append(polka, sign_stack[len(sign_stack)-1])
		sign_stack = sign_stack[:len(sign_stack)-1]
	}
	return polka, nil
}

func getCalc(polka []string) (float64, error) {
	res := make([]float64, 0)
	for _, ch := range polka {
		if isOperation(ch) {
			if len(res) < 2 {
				return float64(0), errors.ErrInvalidExpression
			}
			second := res[len(res)-1]
			first := res[len(res)-2]
			res = res[:len(res)-2]
			resOP, err := applyOperation(first, second, ch)
			if err != nil {
				return float64(0), err
			}
			res = append(res, resOP)
		} else {
			num, err := strconv.ParseFloat(ch, 64)
			if err != nil {
				return float64(0), err
			}
			res = append(res, num)
		}
	}
	if len(res) != 1 {
		return float64(0), errors.ErrInvalidExpression
	}
	return res[0], nil
}

func Calc(expression string) (float64, error) {
	polka, err := parse(expression)
	if err != nil {
		return float64(0), err
	}
	return getCalc(polka)
}
