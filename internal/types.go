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

// FuncMap stores actual fucntion definitions
type FuncMap map[string]CmdDef

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

type InputDocMeta struct {
	InitFuncIsPresent bool
	InitRequired      map[string]struct{} // names of functions required to run init()
}

type InputDoc struct {
	Meta   InputDocMeta
	Tokens []Token
}

func (o *InputDoc) AddInitRequiredFunc(name string) {
	o.Meta.InitRequired[name] = struct{}{}
}

func (o *InputDoc) RemoveInitRequiredFunc(name string) {
	delete(o.Meta.InitRequired, name)
}

func (o *InputDoc) IsInitFuncClarified() bool {
	if o.Meta.InitFuncIsPresent && len(o.Meta.InitRequired) == 0 {
		return true
	}
	return false
}
