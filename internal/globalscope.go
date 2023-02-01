package internal

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
	// fmt.Printf("debug var set: %v:%v\n", key, val)
	vars[key] = val
}

func DeleteVar(key string) {
	// fmt.Printf("debug var del: %v\n", key)
	delete(vars, key)

}

func GetFunc(key string) (*FuncDefinition, bool) {
	// fmt.Printf("debug func get: %v\n", key)
	v, ok := funcs[key]
	return v, ok
}

func SetFunc(key string, val *FuncDefinition) {
	// fmt.Printf("debug func set: %v\n", val)
	funcs[key] = val
}

func DeleteFunc(key string) {
	// fmt.Printf("debug func delete: %v\n", key)
	delete(funcs, key)
}
