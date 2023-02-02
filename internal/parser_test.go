package internal

import (
	"fmt"
	"testing"
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
