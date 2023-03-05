package internal

import "errors"

var (
	MsgBadOperand           = "bad operand %v: %v"
	MsgBadOperandFormat     = "bad operand format %v"
	MsgBadToken             = "bad token [%v: %v]: %v"
	MsgDefaultCmdRun        = "run default cmd %v: %v"
	MsgOperandOutOfRange    = "operand index is out of args range: %v(%v)"
	MsgParsingCmdDefinition = "parsing cmd definition: %v"
	MsgParsingData          = "parsing data: %v"
	MsgPopulateArguments    = "populating arguments: %v"
	MsgUnexpectedCmdKey     = "unexpected cmd key: %v"
	MsgUnexpectedIndex      = "unexpected index value: %v"
	MsgWrongArgReference    = "wrong arg reference %v: %v"

	ErrDivideByZero         = errors.New("divide by zero")
	ErrInitFunctionNotFound = errors.New("init function not found")
	ErrIsNotOrderedMap      = errors.New("function definition is not an (*OrderedMap) type")
	ErrUndefinedFunction    = errors.New("undefined function call")
	ErrUnexpectedCountArgs  = errors.New("unexpected count of arguments")
	ErrUnexpectedValue      = errors.New("unexpected type of value")
	ErrVarIsNotNumber       = errors.New("variable is not a number")
	ErrVarNotFound          = errors.New("var not found")
)
