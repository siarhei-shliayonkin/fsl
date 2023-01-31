package internal

import (
	"errors"
)

var (
	ErrVarNotFound = errors.New("var not found")
)

var vars varMap

func init() {
	vars = make(varMap)
}

func GetVar(key string) (varType, error) {
	v, ok := vars[key]
	if !ok {
		return 0, ErrVarNotFound
	}
	return v, nil
}

func SetVar(key string, val varType) {
	vars[key] = val
}

func DeleteVar(key string) {
	delete(vars, key)
}
