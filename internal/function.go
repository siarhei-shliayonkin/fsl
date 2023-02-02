package internal

// Contains implementation for the Token interface for functions. Also includes a
// code for running default commands and user functions, binding arguments.

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/sirupsen/logrus"
)

var _ Token = (*FuncDefinition)(nil)

func (FuncDefinition) GetType() TokenType            { return TokenTypeFunction }
func (o *FuncDefinition) GetName() string            { return o.Name }
func (o *FuncDefinition) GetDefinition() interface{} { return o }

func (o *FuncDefinition) Print() {
	fmt.Printf("%v\n", o.Name)
	for _, v := range o.Cmds {
		fmt.Printf("  call: %v", v.Call)
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

func (o *FuncDefinition) String() string {
	out := fmt.Sprintf("%v:", o.Name)
	for _, cmd := range o.Cmds {
		out += fmt.Sprintf(" %v", cmd.Call)
	}
	return out
}

// Run runs single cmd or another func
func (o *CmdDef) Run(callArgs ...varType) []string {
	log := logrus.WithField("cmd", o.Call)
	out := []string{}

	cmdArgs, err := o.PopulateArgs(callArgs...)
	if err != nil {
		// log.Errorf("error populating arguments: %v", err)
		// fmt.Printf("undefined\n") // as part of requirements
		out = append(out, "undefined\n")
		return out
	}

	t, isDefault := IsDefaultCmd(o.Call)
	if isDefault {
		msgUnexpectedCoutArgs := "unexpected count of arguments"

		// o.Print()
		switch t {
		case CmdCreateType:
			if len(cmdArgs) != 1 {
				log.Error(msgUnexpectedCoutArgs)
			}
			CmdCreate(o.Target, cmdArgs[0])

		case CmdDeleteType:
			CmdDelete(o.Target)

		case CmdUpdateType:
			if len(cmdArgs) != 1 {
				log.Error(msgUnexpectedCoutArgs)
			}
			CmdUpdate(o.Target, cmdArgs[0])

		case CmdAddType:
			if len(cmdArgs) != 2 {
				log.Error(msgUnexpectedCoutArgs)
			}
			CmdUpdate(o.Target,
				CmdAdd(cmdArgs[0], cmdArgs[1]),
			)

		case CmdSubtractType:
			if len(cmdArgs) != 2 {
				log.Error(msgUnexpectedCoutArgs)
			}
			CmdUpdate(o.Target,
				CmdSubtract(cmdArgs[0], cmdArgs[1]),
			)

		case CmdMultiplyType:
			if len(cmdArgs) != 2 {
				log.Error(msgUnexpectedCoutArgs)
			}
			CmdUpdate(o.Target,
				CmdMultiply(cmdArgs[0], cmdArgs[1]),
			)

		case CmdDivideType:
			if len(cmdArgs) != 2 {
				log.Error(msgUnexpectedCoutArgs)
			}
			CmdUpdate(o.Target,
				CmdDivide(cmdArgs[0], cmdArgs[1]),
			)

		case CmdPrintType:
			if len(cmdArgs) != 1 {
				log.Error(msgUnexpectedCoutArgs)
			}
			out = append(out, CmdPrint(cmdArgs[0]))
		}

		return out
	}

	// another (user) func:
	fn := pureName(o.Call)
	fd, ok := GetFunc(fn)
	if !ok {
		log.Error("undefined function call")
		return out
	}

	for _, cmd := range fd.Cmds {
		if cmd.Target == "$id" {
			cmd.Target = o.Target
		}
		out = append(out, cmd.Run(cmdArgs...)...)
	}
	return out
}

func (o *CmdDef) PopulateArgs(callArgs ...varType) ([]varType, error) {
	cmdArgs := make([]varType, 0, len(o.OperandRefs))

	for _, or := range o.OperandRefs {
		var arg varType

		switch or.ValType {
		case ArgTypeValue:
			arg = or.Value

		case ArgTypeValueRef:
			var err error
			arg, err = GetVar(pureName(or.ValueRef))
			if err != nil {
				return nil, fmt.Errorf("wrong arg reference %v: %v", or.ValueRef, err)
			}

		case ArgTypeOperandRef:
			idx, err := IndexOfOperand(or.ValueRef)
			if err != nil {
				return nil, fmt.Errorf("bad operand %v: %v", or.ValueRef, err)
			}

			logrus.Debugf("operand idx: %v\n", idx)
			if idx < 1 || idx > len(callArgs) {
				return nil, fmt.Errorf("operand index is out of args range: %v(%v)", idx, len(callArgs))
			}

			// arguments counting starts from zero
			idx--
			arg = callArgs[idx]
		}
		cmdArgs = append(cmdArgs, arg)
	}

	return cmdArgs, nil
}

var reIndex = regexp.MustCompile(`^[\$|#][a-z]+`)

func IndexOfOperand(opRef string) (int, error) {
	if !reIndex.MatchString(opRef) {
		return 0, fmt.Errorf("bad operand format %v", opRef)
	}

	parts := reIndex.Split(opRef, 2)

	// if matched then always contains 2 parts
	// if len(parts) != 2 {
	// 	return 0, fmt.Errorf("bad operand reference value %v", opRef)
	// }

	val, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, fmt.Errorf("unexpected index value: %v", opRef)
	}

	return val, nil
}

func (o *CmdDef) Print() {
	fmt.Printf("  call:%v, target:%v, op:", o.Call, o.Target)
	for _, v := range o.OperandRefs {
		fmt.Printf(" %v", v)
	}
	println()
}

// NewFuncToken returns token with function definition
func NewFuncToken(key string, commands []*CmdDef) Token {
	return &FuncDefinition{
		Name: key,
		Args: []varType{},
		Cmds: commands,
	}
}

// CmdCreate adds new variable definition to the global scope
func CmdCreate(name string, value varType) {
	SetVar(name, value)
}

// CmdDelete removes variable definition from the global scope
func CmdDelete(name string) {
	DeleteVar(name)
}

// CmdUpdate updates existing variable definition
func CmdUpdate(name string, value varType) {
	SetVar(name, value)
}

// Add adds two operands
func CmdAdd(arg1, arg2 varType) varType { return arg1 + arg2 }

// Subtract
func CmdSubtract(arg1, arg2 varType) varType { return arg1 - arg2 }

// Multiply
func CmdMultiply(arg1, arg2 varType) varType { return arg1 * arg2 }

// Divide
func CmdDivide(arg1, arg2 varType) varType { return arg1 / arg2 }

// Print
func CmdPrint(arg varType) string { return fmt.Sprintf("%v\n", arg) }

type DefaultCmdType int

const (
	// types of supported default functions
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

// IsDefaultCmd returns type of Cmd and true if the Cmd is default, returns
// false otherwize
func IsDefaultCmd(key string) (DefaultCmdType, bool) {
	t, ok := defaultCmd[key]
	return t, ok
}
