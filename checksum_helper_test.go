package inn

import "testing"

func Test_calculateWeightedSum(t *testing.T) {
	type args struct {
		digits  []int
		weights []int
	}

	tests := []struct {
		name      string
		args      args
		want      int
		wantPanic bool
	}{
		{
			name: "простая сумма",
			args: args{
				digits:  []int{1, 2, 3, 4, 5},
				weights: []int{1, 1, 1, 1, 1},
			},
			want: 15,
		},
		{
			name: "разные веса",
			args: args{
				digits:  []int{1, 2, 3},
				weights: []int{3, 2, 1},
			},
			want: 10, // 1*3 + 2*2 + 3*1 = 10
		},
		{
			name: "с нулями",
			args: args{
				digits:  []int{0, 0, 0},
				weights: []int{5, 10, 3},
			},
			want: 0,
		},
		{
			name: "частично нули",
			args: args{
				digits:  []int{0, 2, 0},
				weights: []int{5, 10, 3},
			},
			want: 20, // 2*10
		},
		{
			name: "отрицательные значения",
			args: args{
				digits:  []int{-1, 2},
				weights: []int{3, 4},
			},
			want: 5, // -1*3 + 2*4 = 5
		},
		{
			name: "пустые срезы",
			args: args{
				digits:  []int{},
				weights: []int{},
			},
			want: 0,
		},
		{
			name: "разная длина -> panic",
			args: args{
				digits:  []int{1, 2},
				weights: []int{1},
			},
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("expected panic but did not panic")
					}
				}()
			}

			got := calculateWeightedSum(tt.args.digits, tt.args.weights)

			if !tt.wantPanic && got != tt.want {
				t.Errorf("calculateWeightedSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateChecksumDigit(t *testing.T) {
	type args struct {
		weightedSum int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "остаток меньше 10",
			args: args{weightedSum: 22}, // 22 % 11 = 0
			want: 0,
		},
		{
			name: "остаток 1",
			args: args{weightedSum: 23}, // 23 % 11 = 1
			want: 1,
		},
		{
			name: "остаток 9",
			args: args{weightedSum: 20}, // 20 % 11 = 9
			want: 9,
		},
		{
			name: "остаток 10 -> редукция",
			args: args{weightedSum: 21}, // 21 % 11 = 10 -> 10 % 10 = 0
			want: 0,
		},
		{
			name: "большое число с редукцией",
			args: args{weightedSum: 32}, // 32 % 11 = 10 -> 0
			want: 0,
		},
		{
			name: "большое число без редукции",
			args: args{weightedSum: 33}, // 33 % 11 = 0
			want: 0,
		},
		{
			name: "еще редукция",
			args: args{weightedSum: 43}, // 43 % 11 = 10 -> 0
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateChecksumDigit(tt.args.weightedSum); got != tt.want {
				t.Errorf("calculateChecksumDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}
