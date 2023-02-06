package internal

import (
	"encoding/json"
	"testing"
)

func Test_getVarValue(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    varType
		wantErr bool
	}{
		{
			name: "1, valid",
			args: args{
				value: json.Number("1"),
			},
			want:    varType(1),
			wantErr: false,
		},
		{
			name: "1.1, bad",
			args: args{
				value: 1,
			},
			want:    varType(0),
			wantErr: true,
		},
	}
	vd := &VariableDefinition{
		Name:  "id",
		Value: varType(1),
	}
	_ = vd.GetType()
	_ = vd.GetName()
	_ = vd.GetDefinition()
	vd.Print()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getVarValue(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("getVarValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getVarValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
