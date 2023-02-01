package internal

/*
  Represents global storage for variable and function definitions.
*/

import (
	"errors"
)

// FuncMap stores actual fucntion definitions
type FuncMap map[string]*FuncDefinition

var (
	ErrVarNotFound = errors.New("var not found")
)

var (
	vars  varMap
	funcs FuncMap
)

func init() {
	vars = make(varMap)
	funcs = make(map[string]*FuncDefinition)
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

func GetFunc(key string) (*FuncDefinition, bool) {
	v, ok := funcs[key]
	return v, ok
}

func SetFunc(key string, val *FuncDefinition) {
	funcs[key] = val
}

func DeleteFunc(key string) {
	delete(funcs, key)
}
