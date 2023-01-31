package internal

import "fmt"

var _ Token = (*FuncDefinition)(nil)

func (FuncDefinition) GetType() TokenType            { return TokenTypeFunction }
func (o *FuncDefinition) GetName() string            { return o.Name }
func (o *FuncDefinition) GetDefinition() interface{} { return o }
func (o *FuncDefinition) Print() {
	fmt.Printf("%v\n", o.Name)
	for _, v := range o.Cmds {
		fmt.Printf("  call: %v", v.Call) //, v.Target, v.OperandRefs)
		if len(v.Target) > 0 {
			fmt.Printf(", target: %v", v.Target)
		}
		if len(v.OperandRefs) > 0 {
			for i, op := range v.OperandRefs {
				fmt.Printf(", op%d: %v", i, op)
			}
		}
		println()
	}
}

func NewFuncToken(key string, commands []*CmdDef) Token {
	return &FuncDefinition{
		Name: key,
		Args: []varType{},
		Cmds: commands,
	}
}

// list of supported functions
// - create
// - delete
// - update
// - add
// - subtract
// - multiply
// - divide
// - print

// Create adds new variable to the global scope definition
func CmdCreate(name string, value varType) {
	SetVar(name, value)
}

// Delete ..
func CmdDelete(name string) {
	DeleteVar(name)
}

// Update ..
func CmdUpdate(name string, value varType) {
	SetVar(name, value)
}

// Add ..
func CmdAdd(arg1, arg2 varType) varType { return arg1 + arg2 }

// Subtract ..
func CmdSubtract(arg1, arg2 varType) varType { return arg1 - arg2 }

// Multiply ..
func CmdMultiply(arg1, arg2 varType) varType { return arg1 * arg2 }

// Divide ..
func CmdDivide(arg1, arg2 varType) varType { return arg1 / arg2 }

// Print ..
func CmdPrint(arg varType) { fmt.Printf("%v\n", arg) }

// func NewOperatorDefinition(cmd string, target string, operands []string) *OperatorType {}

// func (c *OperatorType) runCmd(args []varType) {
// 	switch c.Cmd {
// 	case "print":
// 		if len(args) == 0 {
// 			fmt.Printf("no arguments provided")
// 			return
// 		}
// 		CmdPrint(args[0])
// 	default:
// 		fmt.Printf("operation %v is not supported", c.Cmd)
// 	}
// }

// func (f *FuncDefinition) runFunc(operands []ArgumentDefinition) {
// 	for _, op := range f.Operators {
// 		// fmt.Printf("%v\n", op)

// 		// TODO: args binding

// 		op.runCmd([]varType{123})
// 	}
// }

type DefaultCmdType int

const (
	CmdCreateType DefaultCmdType = iota
	CmdDeleteType
	CmdUpdateType
	CmdAddType
	CmdSubtractType
	CmdMultiplyType
	CmdDivideType
	CmdPrintType
)

var defaultCmd map[string]DefaultCmdType = map[string]DefaultCmdType{
	"create":   CmdCreateType,
	"delete":   CmdDeleteType,
	"update":   CmdUpdateType,
	"add":      CmdAddType,
	"subtract": CmdSubtractType,
	"multiply": CmdMultiplyType,
	"divide":   CmdDivideType,
	"print":    CmdPrintType,
}

func IsDefaultCmd(key string) bool {
	if _, ok := defaultCmd[key]; ok {
		return true
	}
	return false
}

func (d *InputDoc) PrintDoc() {
	// meta
	println("-- new doc --")
	println("Meta:")
	println("  InitFuncIsPresent:", d.Meta.InitFuncIsPresent)
	println("  InitRequired:")
	for k := range d.Meta.InitRequired {
		println("   ", k)
	}

	// tokens
	println("--")
	for _, v := range d.Tokens {
		v.Print()
	}
}
