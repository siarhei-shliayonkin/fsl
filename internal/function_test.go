package internal

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexOfOperand(t *testing.T) {
	type args struct {
		opRef string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "1, valid",
			args: args{
				opRef: "$value1",
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "2, bad value1",
			args: args{
				opRef: "bad value1",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "3",
			args: args{
				opRef: "#var3",
			},
			want:    3,
			wantErr: false,
		},
		{
			name: "4, not valid",
			args: args{
				opRef: "$value1_",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "5, not valid, more then 2 parts",
			args: args{
				opRef: "$value1nextpart",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "6, not valid, no index",
			args: args{
				opRef: "$value",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "7, not valid, starts with no special symbol",
			args: args{
				opRef: "value",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := indexOfOperand(tt.args.opRef)
			if (err != nil) != tt.wantErr {
				t.Errorf("indexOfOperand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("indexOfOperand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFuncDefinition_String(t *testing.T) {
	type fields struct {
		Name string
		Args []varType
		Cmds []*CmdDef
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "1",
			fields: fields{
				Name: "1",
				Args: []varType{},
				Cmds: []*CmdDef{
					{
						Call:        "1",
						Target:      "1",
						OperandRefs: []*ArgDef{},
					},
				},
			},
			want: "1: 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &FuncDefinition{
				Name: tt.fields.Name,
				Args: tt.fields.Args,
				Cmds: tt.fields.Cmds,
			}
			if got := o.String(); got != tt.want {
				t.Errorf("FuncDefinition.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCmdDef_PopulateArgs(t *testing.T) {
	vars = make(varMap)
	SetVar("var1", 1)

	type fields struct {
		Call        string
		Target      string
		OperandRefs []*ArgDef
	}
	type args struct {
		callArgs []varType
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []varType
		wantErr bool
	}{
		{
			name: "1, value",
			fields: fields{
				Call:        "print",
				Target:      "",
				OperandRefs: []*ArgDef{NewArgValue(11)},
			},
			args: args{
				callArgs: []varType{},
			},
			want:    []varType{11},
			wantErr: false,
		},
		{
			name: "2, ValueRef",
			fields: fields{
				Call:        "print",
				Target:      "",
				OperandRefs: []*ArgDef{NewArgValueRef("#var1")},
			},
			args: args{
				callArgs: []varType{},
			},
			want:    []varType{1},
			wantErr: false,
		},
		{
			name: "2.1, bad ValueRef",
			fields: fields{
				Call:        "print",
				Target:      "",
				OperandRefs: []*ArgDef{NewArgValueRef("#badvar1")},
			},
			args: args{
				callArgs: []varType{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "3, OperandRef",
			fields: fields{
				Call:        "print",
				Target:      "",
				OperandRefs: []*ArgDef{NewArgOperandRef("$value1")},
			},
			args: args{
				callArgs: []varType{3.3},
			},
			want:    []varType{3.3},
			wantErr: false,
		},
		{
			name: "3.1, bad OperandRef",
			fields: fields{
				Call:        "print",
				Target:      "",
				OperandRefs: []*ArgDef{NewArgOperandRef("badvalue1")},
			},
			args: args{
				callArgs: []varType{3.3},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "3.2, bad OperandRef index",
			fields: fields{
				Call:        "print",
				Target:      "",
				OperandRefs: []*ArgDef{NewArgOperandRef("$value2")},
			},
			args: args{
				callArgs: []varType{3.3},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &CmdDef{
				Call:        tt.fields.Call,
				Target:      tt.fields.Target,
				OperandRefs: tt.fields.OperandRefs,
			}
			got, err := o.populateArgs(tt.args.callArgs...)
			if (err != nil) != tt.wantErr {
				t.Errorf("CmdDef.populateArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CmdDef.populateArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCmdDef_runDefaultCmd(t *testing.T) {
	cmd := NewCmdDef()
	checkVal := func(valRef string, valExpected varType) {
		valActual, err := GetVar(valRef)
		assert.NoError(t, err)
		assert.Equal(t, valExpected, valActual)
	}

	// CmdCreateType
	cmd.Target = "var1"
	valExpected := varType(1)
	_, err := cmd.runDefaultCmd(CmdCreateType, []varType{valExpected}...)
	assert.NoError(t, err)
	checkVal(cmd.Target, valExpected)
	// bad args
	_, err = cmd.runDefaultCmd(CmdCreateType)
	assert.Error(t, err)

	// CmdDeleteType
	_, err = cmd.runDefaultCmd(CmdDeleteType)
	assert.NoError(t, err)
	_, err = GetVar(cmd.Target)
	assert.Error(t, err, "does not exist anymore")

	// CmdUpdateType: var1=1
	_, err = cmd.runDefaultCmd(CmdUpdateType, []varType{valExpected}...)
	assert.NoError(t, err)
	checkVal(cmd.Target, valExpected)
	_, err = cmd.runDefaultCmd(CmdUpdateType)
	assert.Error(t, err)

	// CmdAddType: var1=7
	_, err = cmd.runDefaultCmd(CmdAddType, []varType{varType(3), varType(4)}...)
	assert.NoError(t, err)
	checkVal(cmd.Target, varType(7))
	_, err = cmd.runDefaultCmd(CmdAddType, []varType{varType(3), varType(4), varType(5)}...)
	assert.Error(t, err)

	// CmdSubtractType
	_, err = cmd.runDefaultCmd(CmdSubtractType, []varType{varType(1), varType(4.5)}...)
	assert.NoError(t, err)
	checkVal(cmd.Target, varType(-3.5))
	_, err = cmd.runDefaultCmd(CmdSubtractType)
	assert.Error(t, err)

	// CmdMultiplyType
	_, err = cmd.runDefaultCmd(CmdMultiplyType, []varType{varType(-2), varType(-3.5)}...)
	assert.NoError(t, err)
	checkVal(cmd.Target, varType(7))
	_, err = cmd.runDefaultCmd(CmdMultiplyType)
	assert.Error(t, err)

	// CmdDivideType
	_, err = cmd.runDefaultCmd(CmdDivideType, []varType{varType(15), varType(5)}...)
	assert.NoError(t, err)
	checkVal(cmd.Target, varType(3))
	_, err = cmd.runDefaultCmd(CmdDivideType, []varType{varType(15), varType(0)}...)
	assert.Error(t, err, "divide by zero")
	checkVal(cmd.Target, varType(3))
	_, err = cmd.runDefaultCmd(CmdDivideType)
	assert.Error(t, err)

	// CmdPrintType
	_, err = cmd.runDefaultCmd(CmdPrintType, []varType{varType(11)}...)
	assert.NoError(t, err)
	_, err = cmd.runDefaultCmd(CmdPrintType)
	assert.Error(t, err)
}

func TestCmdDef_Run(t *testing.T) {
	// define  "sum" func
	cmd := NewCmdDef()
	cmd.Call = "add"
	cmd.Target = "$id"
	cmd.OperandRefs = append(cmd.OperandRefs, []*ArgDef{
		NewArgOperandRef("$value1"),
		NewArgOperandRef("$value2"),
	}...)

	fd := &FuncDefinition{
		Name: "sum",
		Cmds: []*CmdDef{cmd},
	}
	SetFunc("sum", fd)
	_ = fd.GetType()
	_ = fd.GetName()
	_ = fd.GetDefinition()
	_ = NewFuncToken("sum", fd.Cmds)

	type fields struct {
		Call        string
		Target      string
		OperandRefs []*ArgDef
	}
	type args struct {
		callArgs []varType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			name: "1, add",
			fields: fields{
				Call:   "add",
				Target: "var1",
				OperandRefs: []*ArgDef{
					NewArgValue(1),
					NewArgValue(2),
				},
			},
			args: args{
				callArgs: []varType{},
			},
			want: []string{},
		},
		{
			name: "1.1, add check",
			fields: fields{
				Call:   "print",
				Target: "",
				OperandRefs: []*ArgDef{
					NewArgValueRef("var1"),
				},
			},
			args: args{
				callArgs: []varType{},
			},
			want: []string{"3"},
		},
		{
			name: "1.2, add, bad argRef",
			fields: fields{
				Call:   "add",
				Target: "var1",
				OperandRefs: []*ArgDef{
					NewArgValueRef("var2"),
				},
			},
			args: args{
				callArgs: []varType{},
			},
			want: []string{"undefined"},
		},
		{
			name: "1.3, add, bad args count",
			fields: fields{
				Call:   "add",
				Target: "var1",
				OperandRefs: []*ArgDef{
					NewArgValue(1),
				},
			},
			args: args{
				callArgs: []varType{},
			},
			want: []string{},
		},
		{
			name: "2, sum, user func",
			fields: fields{
				Call:   "#sum",
				Target: "var1",
				OperandRefs: []*ArgDef{
					NewArgValueRef("var1"),
					NewArgValue(3.5),
				},
			},
			args: args{
				callArgs: []varType{},
			},
			want: []string{},
		},
		{
			name: "2.1, sum check",
			fields: fields{
				Call:   "print",
				Target: "",
				OperandRefs: []*ArgDef{
					NewArgValueRef("var1"),
				},
			},
			args: args{
				callArgs: []varType{},
			},
			want: []string{"6.5"},
		},
		{
			name: "3, undefined func",
			fields: fields{
				Call:   "#someFunc",
				Target: "",
				OperandRefs: []*ArgDef{
					NewArgValueRef("var1"),
				},
			},
			args: args{
				callArgs: []varType{},
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &CmdDef{
				Call:        tt.fields.Call,
				Target:      tt.fields.Target,
				OperandRefs: tt.fields.OperandRefs,
			}
			if got := o.Run(tt.args.callArgs...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CmdDef.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
