package inn

import "testing"

func Test_isDigit(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "число",
			args: args{
				r: '5',
			},
			want: true,
		},
		{
			name: "граничный случай 0",
			args: args{
				r: '0',
			},
			want: true,
		},
		{
			name: "граничный случай 9",
			args: args{
				r: '9',
			},
			want: true,
		},
		{
			name: "не число",
			args: args{
				r: 'z',
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isDigit(tt.args.r); got != tt.want {
				t.Errorf("isDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}
