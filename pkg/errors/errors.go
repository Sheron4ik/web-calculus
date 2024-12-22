package errors

import "errors"

var (
	ErrMissingOpenBracket  = errors.New("missing `(` in expression")
	ErrMissingCloseBracket = errors.New("missing `)` in expression")
	ErrDivideByZero        = errors.New("division by zero")
	ErrUnknownOperations   = errors.New("unknown operation in expression")
	ErrInvalidExpression   = errors.New("invalid expression")
	ErrEmptyExpression     = errors.New("empty expression")
)
