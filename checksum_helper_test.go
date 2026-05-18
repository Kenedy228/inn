package inn

import "testing"

func Test_calculateWeightedSum(t *testing.T) {
	type args struct {
		digits  []int
		weights []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "простая сумма",
			args: args{
				digits:  []int{1, 2, 3, 4, 5},
				weights: []int{1, 1, 1, 1, 1},
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateWeightedSum(tt.args.digits, tt.args.weights); got != tt.want {
				t.Errorf("calculateWeightedSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
