package calculus

// import (
// 	"testing"

// 	"github.com/Sheron4ik/web-calculus/pkg/errors"
// )

// func TestCalc(t *testing.T) {
// 	tests := []struct {
// 		name            string
// 		expression      string
// 		expected_result float64
// 		expected_error  error
// 	}{
// 		{
// 			name:            "simple addition",
// 			expression:      "1+2",
// 			expected_result: 3.0,
// 			expected_error:  nil,
// 		},
// 		{
// 			name:            "simple subtraction",
// 			expression:      "4-3",
// 			expected_result: 1.0,
// 			expected_error:  nil,
// 		},
// 		{
// 			name:            "simple multiplication",
// 			expression:      "5*6",
// 			expected_result: 30.0,
// 			expected_error:  nil,
// 		},
// 		{
// 			name:            "simple division",
// 			expression:      "7/8",
// 			expected_result: 0.875,
// 			expected_error:  nil,
// 		},
// 		{
// 			name:            "brackets",
// 			expression:      "(0+9)",
// 			expected_result: 9.0,
// 			expected_error:  nil,
// 		},
// 		{
// 			name:            "complex expression",
// 			expression:      "2*(1+3)/4",
// 			expected_result: 2.0,
// 			expected_error:  nil,
// 		},
// 		{
// 			name:            "priority expression",
// 			expression:      "2+2*2",
// 			expected_result: 6.0,
// 			expected_error:  nil,
// 		},
// 		{
// 			name:            "missing open bracket",
// 			expression:      "0+0)",
// 			expected_result: 0.0,
// 			expected_error:  errors.ErrMissingOpenBracket,
// 		},
// 		{
// 			name:            "missing close bracket",
// 			expression:      "(0+0",
// 			expected_result: 0.0,
// 			expected_error:  errors.ErrMissingCloseBracket,
// 		},
// 		{
// 			name:            "division by zero",
// 			expression:      "0/0",
// 			expected_result: 0.0,
// 			expected_error:  errors.ErrDivideByZero,
// 		},
// 		{
// 			name:            "unknown operation",
// 			expression:      "1+2^3",
// 			expected_result: 0.0,
// 			expected_error:  errors.ErrUnknownOperations,
// 		},
// 		{
// 			name:            "invalid expression",
// 			expression:      "1+2+",
// 			expected_result: 0.0,
// 			expected_error:  errors.ErrInvalidExpression,
// 		},
// 		{
// 			name:            "empty expression",
// 			expression:      "",
// 			expected_result: 0.0,
// 			expected_error:  errors.ErrEmptyExpression,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			result, err := Calc(test.expression)

// 			if err != test.expected_error {
// 				t.Errorf("Calc(%s): expected error: %s, got: %s", test.expression, test.expected_error.Error(), err.Error())
// 			}

// 			if result != test.expected_result {
// 				t.Errorf("Calc(%s): expected result: %f, got: %f", test.expression, test.expected_result, result)
// 			}
// 		})
// 	}
// }
