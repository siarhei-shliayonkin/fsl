package internal

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
