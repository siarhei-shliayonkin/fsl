package internal

import "testing"

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
