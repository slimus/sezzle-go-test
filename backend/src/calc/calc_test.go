package calc

import "testing"

func TestCalculate(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "correct sum",
			args: args{
				expression: "3+1",
			},
			want:    4,
			wantErr: false,
		},
		{
			name: "sum with error",
			args: args{
				expression: "a+1",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "sum with error",
			args: args{
				expression: "1+b",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "sum with error",
			args: args{
				expression: "1.42b+2.23",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "correct sub",
			args: args{
				expression: "1-10",
			},
			want:    -9,
			wantErr: false,
		},
		{
			name: "correct sub",
			args: args{
				expression: "0-0",
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "correct multiply",
			args: args{
				expression: "100*500",
			},
			want:    50000,
			wantErr: false,
		},
		{
			name: "correct multiply",
			args: args{
				expression: "4.2*10",
			},
			want:    42,
			wantErr: false,
		},
		{
			name: "correct dev",
			args: args{
				expression: "10/4",
			},
			want:    2.5,
			wantErr: false,
		},
		{
			name: "bad expression",
			args: args{
				expression: "a+",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "bad expression",
			args: args{
				expression: "*a",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Calculate(tt.args.expression)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calculate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
