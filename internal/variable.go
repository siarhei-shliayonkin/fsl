package internal

// Contains implementation for the Token interface for variable definition.

import (
	"encoding/json"
	"fmt"
)

func NewVarToken(id string, value interface{}) (Token, error) {
	val, err := getVarValue(value)
	if err != nil {
		return nil, err
	}
	return &VariableDefinition{
		Name:  id,
		Value: val,
	}, nil
}

var _ Token = (*VariableDefinition)(nil)

func (VariableDefinition) GetType() TokenType            { return TokenTypeVariable }
func (o *VariableDefinition) GetName() string            { return o.Name }
func (o *VariableDefinition) GetDefinition() interface{} { return o }
func (o *VariableDefinition) Print() {
	fmt.Printf("%v:%v\n", o.Name, o.Value)
}

func getVarValue(value interface{}) (varType, error) {
	pairVal, ok := value.(json.Number)
	if !ok {
		return 0, ErrVarIsNotNumber
	}

	v, err := pairVal.Float64()
	if err != nil {
		return 0, ErrVarIsNotNumber
	}
	return varType(v), nil
}
