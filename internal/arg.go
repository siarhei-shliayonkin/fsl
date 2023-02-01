package internal

// Contains description of various types of argument definition

import "fmt"

type ArgType int

const (
	ArgTypeValue ArgType = iota + 1
	ArgTypeValueRef
	ArgTypeOperandRef
)

type ArgDef struct {
	ValType  ArgType
	ValueRef string
	Value    varType
}

func NewArgValueRef(ref string) *ArgDef {
	return &ArgDef{
		ValType:  ArgTypeValueRef,
		ValueRef: ref,
	}
}
func NewArgValue(val varType) *ArgDef {
	return &ArgDef{
		ValType: ArgTypeValue,
		Value:   val,
	}
}
func NewArgOperandRef(ref string) *ArgDef {
	return &ArgDef{
		ValType:  ArgTypeOperandRef,
		ValueRef: ref,
	}
}

func (o ArgDef) String() string {
	var out string
	switch o.ValType {
	case ArgTypeValue:
		out = fmt.Sprintf("%v", o.Value)
	case ArgTypeValueRef, ArgTypeOperandRef:
		out = o.ValueRef
	}
	return out
}
