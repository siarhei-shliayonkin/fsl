package internal

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	ojson "gitlab.com/c0b/go-ordered-json"
)

func Test_typeOfKey(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want CmdKeyType
	}{
		{
			name: "cmd",
			args: args{
				key: "cmd",
			},
			want: CmdKeyTypeCmd,
		},
		{
			name: "id",
			args: args{
				key: "id",
			},
			want: CmdKeyTypeID,
		},
		{
			name: "value",
			args: args{
				key: "value",
			},
			want: CmdKeyTypeValue,
		},
		{
			name: "value1",
			args: args{
				key: "value1",
			},
			want: CmdKeyTypeValue,
		},
		{
			name: "operand1",
			args: args{
				key: "operand1",
			},
			want: CmdKeyTypeOperand,
		},
		{
			name: "undefined",
			args: args{
				key: "undefined",
			},
			want: CmdKeyTypeUndefined,
		},
	}

	fmt.Printf("CmdKeyTypeCmd:%v, CmdKeyTypeID:%v\n", CmdKeyTypeCmd, CmdKeyTypeID)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := typeOfKey(tt.args.key); got != tt.want {
				t.Errorf("typeOfKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseInput(t *testing.T) {
	tokenForVar := func(key, val string) Token {
		token, err := NewVarToken(key, json.Number(val))
		assert.NoError(t, err)
		return token
	}

	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *Script
		wantErr bool
	}{
		{
			name: "1, empty",
			args: args{
				data: []byte("{}"),
			},
			want:    NewScript(),
			wantErr: false,
		},
		{
			name: "1, bad",
			args: args{
				data: []byte(""),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "2, vars",
			args: args{
				data: []byte(`{"var1":11, "var2":12}`),
			},
			want: &Script{
				Meta: ScriptMeta{},
				Tokens: []Token{
					tokenForVar("var1", "11"),
					tokenForVar("var2", "12"),
				},
				Output: []string{},
			},
			wantErr: false,
		},
		{
			name: "2, func",
			args: args{
				data: []byte(`{"init": [ {"cmd" : "#setup" } ]}`),
			},
			want: &Script{
				Meta: ScriptMeta{},
				Tokens: []Token{
					NewFuncToken("init", []*CmdDef{{
						Call:        "#setup",
						Target:      "",
						OperandRefs: []*ArgDef{},
					}}),
				},
			},
			wantErr: false,
		},
		{
			name: "3, func, bad cmd",
			args: args{
				data: []byte(`{"init": [ {"cmd":"print", "val": "#var1"} ]}`),
			},
			want: &Script{
				Meta: ScriptMeta{},
				Tokens: []Token{
					NewFuncToken("init", []*CmdDef{{
						Call:        "print",
						Target:      "",
						OperandRefs: []*ArgDef{},
					}}),
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseInput(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				fmt.Println(err)
				return
			}
			if !assert.Equal(t, tt.want.Tokens, got.Tokens) {
				t.Errorf("ParseInput() = %v, want %v", got.Tokens, tt.want.Tokens)
			}
		})
	}
}

func Test_parseCmd(t *testing.T) {
	tests := []struct {
		name    string
		sample  *ojson.OrderedMap
		want    *CmdDef
		wantErr bool
	}{
		{
			name: "1, by valueRef",
			sample: ojson.NewOrderedMapFromKVPairs([]*ojson.KVPair{
				{Key: "cmd", Value: "print"},
				{Key: "value1", Value: "#var1"},
			}),
			want: &CmdDef{
				Call:        "print",
				Target:      "",
				OperandRefs: []*ArgDef{NewArgValueRef("#var1")},
			},
			wantErr: false,
		},
		{
			name: "2, by value",
			sample: ojson.NewOrderedMapFromKVPairs([]*ojson.KVPair{
				{Key: "cmd", Value: "print"},
				{Key: "value1", Value: json.Number("11")},
			}),
			want: &CmdDef{
				Call:        "print",
				Target:      "",
				OperandRefs: []*ArgDef{NewArgValue(11)},
			},
			wantErr: false,
		},
		{
			name: "2.1, by value",
			sample: ojson.NewOrderedMapFromKVPairs([]*ojson.KVPair{
				{Key: "cmd", Value: "print"},
				{Key: "value1", Value: json.Number("ab")},
			}),
			want:    nil,
			wantErr: true,
		},
		{
			name: "3, by operand",
			sample: ojson.NewOrderedMapFromKVPairs([]*ojson.KVPair{
				{Key: "cmd", Value: "print"},
				{Key: "operand1", Value: "$value1"},
			}),
			want: &CmdDef{
				Call:        "print",
				Target:      "",
				OperandRefs: []*ArgDef{NewArgOperandRef("$value1")},
			},
			wantErr: false,
		},
		{
			name: "3.1, by operand, bad operand",
			sample: ojson.NewOrderedMapFromKVPairs([]*ojson.KVPair{
				{Key: "cmd", Value: "print"},
				{Key: "operand1", Value: 11},
			}),
			want:    nil,
			wantErr: true,
		},
		{
			name: "4, id",
			sample: ojson.NewOrderedMapFromKVPairs([]*ojson.KVPair{
				{Key: "cmd", Value: "add"},
				{Key: "id", Value: "$id"},
				{Key: "value1", Value: json.Number("1")},
				{Key: "value2", Value: json.Number("2")},
			}),
			want: &CmdDef{
				Call:        "add",
				Target:      "$id",
				OperandRefs: []*ArgDef{NewArgValue(1), NewArgValue(2)},
			},
			wantErr: false,
		},
		{
			name: "4.1, id, bad",
			sample: ojson.NewOrderedMapFromKVPairs([]*ojson.KVPair{
				{Key: "cmd", Value: "add"},
				{Key: "id", Value: 1},
				{Key: "value1", Value: json.Number("1")},
				{Key: "value2", Value: json.Number("2")},
			}),
			want:    nil,
			wantErr: true,
		},
		{
			name: "5.1, cmd, bad",
			sample: ojson.NewOrderedMapFromKVPairs([]*ojson.KVPair{
				{Key: "cmd", Value: 1},
				{Key: "id", Value: 1},
				{Key: "value1", Value: json.Number("1")},
				{Key: "value2", Value: json.Number("2")},
			}),
			want:    nil,
			wantErr: true,
		},
		{
			name: "5.2, cmd, no cmd",
			sample: ojson.NewOrderedMapFromKVPairs([]*ojson.KVPair{
				{Key: "cmd1", Value: 1},
				{Key: "id", Value: 1},
				{Key: "value1", Value: json.Number("1")},
				{Key: "value2", Value: json.Number("2")},
			}),
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseCmd(tt.sample.EntriesIter())
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCmd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseCmd() = %v, want %v", got, tt.want)
			}
		})
	}
}
