package internal

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// Processing the document created from the input data

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

func (o *InputDoc) PrintDoc() {
	// meta
	println("-- new doc --")
	println("Meta:")
	println("  InitFuncIsPresent:", o.Meta.InitFuncIsPresent)
	println("  InitRequired:")
	for k := range o.Meta.InitRequired {
		println("   ", k)
	}

	// tokens
	println("--")
	for _, v := range o.Tokens {
		v.Print()
	}
}

func (o *InputDoc) Process() {
	for _, token := range o.Tokens {
		switch token.GetType() {
		case TokenTypeVariable:
			vd, ok := token.GetDefinition().(*VariableDefinition)
			if !ok {
				fmt.Printf("err: token.GetDefinition() for variable %v\n", token.GetName())
			}
			SetVar(token.GetName(), vd.Value)

		case TokenTypeFunction:
			updateGlobalFuncDefinition(token)
			if o.IsInitRequired(token.GetName()) {
				o.RemoveFromInitRequired(token.GetName())
			}
		}
	}

	if o.IsInitResolved() {
		runInitFunc()
		o.Meta.InitFuncIsPresent = false // prevent next init() call, if new doc doesn't contain init()
	}
}

func (o *InputDoc) RemoveFromInitRequired(name string) {
	delete(o.Meta.InitRequired, name)
}

func (o *InputDoc) IsInitRequired(name string) bool {
	_, ok := o.Meta.InitRequired[name]
	return ok
}

func (o *InputDoc) IsInitResolved() bool {
	return o.Meta.InitFuncIsPresent && len(o.Meta.InitRequired) == 0
}

func updateGlobalFuncDefinition(token Token) {
	vd, ok := token.GetDefinition().(*FuncDefinition)
	if !ok {
		fmt.Printf("err: token.GetDefinition() for func %v\n", token.GetName())
		return
	}
	SetFunc(token.GetName(), vd)
}

func runInitFunc() {
	defer DeleteFunc("init")

	fd, ok := GetFunc("init")
	if !ok {
		fmt.Printf("err: init function not found")
	}

	for _, cmd := range fd.Cmds {
		logrus.Debugf("run init, cmd: %v", cmd)
		cmd.Run()
	}
}
