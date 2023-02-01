package internal

type varType float64           // default type for numeric values
type varMap map[string]varType // storage for all actual variables

// VariableDefinition ..
type VariableDefinition struct {
	Name  string
	Value varType
}

// CmdDef ..
type CmdDef struct {
	Call        string
	Target      string
	OperandRefs []*ArgDef
}

// FuncDefinition describes function definition
type FuncDefinition struct {
	Name string
	Args []varType // actual values, populated on func call
	Cmds []*CmdDef
}

// OperatorType ..
// type OperatorType struct {
// 	Cmd      string
// 	Target   string
// 	Operands []string
// }

// ArgumentDefinition describes argument to be resolved to values on call
// type ArgumentDefinition struct {
// 	ArgType string // TODO: const
// 	ArgRef  string // format: "#var1", "$value2"
// }
