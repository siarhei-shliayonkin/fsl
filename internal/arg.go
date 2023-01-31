package internal

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

// func NewArgOperandRef() ArgOperandRef { return "" }

// func (o ArgDef) Populate() varType {
// 	// TODO: switch o.ValType
// 	return 1
// }

// type ArgValue interface {
// 	Populate() varType
// }

// type ArgValueRef string
// type ArgOperandRef string
