package internal

// Contains implementation for the Token interface for variable definition.

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

func NewVarToken(id string, value interface{}) Token {
	val, err := getVarValue(value)
	if err != nil {
		log.Fatalln(err)
	}
	return &VariableDefinition{
		Name:  id,
		Value: val,
	}
}

var _ Token = (*VariableDefinition)(nil)

func (VariableDefinition) GetType() TokenType            { return TokenTypeVariable }
func (o *VariableDefinition) GetName() string            { return o.Name }
func (o *VariableDefinition) GetDefinition() interface{} { return o }
func (o *VariableDefinition) Print() {
	fmt.Printf("%v:%v\n", o.Name, o.Value)
}

var ErrVarIsNotNumber = errors.New("variable is not a number")

func getVarValue(value interface{}) (varType, error) {
	pairVal, ok := value.(json.Number)
	if !ok {
		return 0, ErrVarIsNotNumber
	}

	v, err := pairVal.Float64()
	if err != nil {
		return 0, ErrVarIsNotNumber
	}

	return varType(v), err
}
