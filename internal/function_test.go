package internal

import "testing"

// func _1_TestFuncDefinition_runFunc(t *testing.T) {
// 	type fields struct {
// 		Name      string
// 		Args      []varType
// 		Operators []*OperatorType
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 	}{
// 		{
// 			name: "1",
// 			fields: fields{
// 				Name: "",
// 				Args: []varType{1},
// 				Operators: []*OperatorType{
// 					{
// 						Cmd:    "print",
// 						Target: "",
// 						Operands: []string{
// 							`"value": "#var1"`,
// 						},
// 					},
// 				},
// 			},
// 		},
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			f := FuncDefinition{
// 				Name:      tt.fields.Name,
// 				Args:      tt.fields.Args,
// 				Operators: tt.fields.Operators,
// 			}
// 			f.runFunc()
// 		})
// 	}
// }

// func TestFuncDefinition_runFunc(t *testing.T) {
// 	type fields struct {
// 		Name      string
// 		Args      []varType
// 		Operators []*OperatorType
// 	}
// 	type args struct {
// 		operands []ArgumentDefinition
// 	}
// tests := []struct {
// 	name   string
// 	fields fields
// 	args   args
// }{
// 	{
// 		name:   "1",
// 		fields: fields{Name: "", Args: []varType{1}, Operators: []*OperatorType{{Cmd: "print", Target: "", Operands: []string{`"value": "#var1"`}}}},
// 		args: args{
// 			operands: []ArgumentDefinition{
// 				ArgumentDefinition{
// 					ArgType: "",
// 					ArgRef:  "",
// 				},
// 			},
// 		},
// 	},

// 	// TODO: Add test cases.
// }

// for _, tt := range tests {
// 	t.Run(tt.name, func(t *testing.T) {
// 		f := &FuncDefinition{
// 			Name:      tt.fields.Name,
// 			Args:      tt.fields.Args,
// 			Operators: tt.fields.Operators,
// 		}
// 		f.runFunc(tt.args.operands)
// 	})
// }
// }

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
			got, err := IndexOfOperand(tt.args.opRef)
			if (err != nil) != tt.wantErr {
				t.Errorf("IndexOfOperand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IndexOfOperand() = %v, want %v", got, tt.want)
			}
		})
	}
}
