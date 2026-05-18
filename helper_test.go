package inn

import (
	"reflect"
	"strings"
	"testing"
)

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

func Test_parseDigits(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name      string
		args      args
		want      []int
		wantPanic bool
	}{
		{
			name: "парсинг строки из всех нулей",
			args: args{
				s: strings.Repeat("0", 20),
			},
			want:      make([]int, 20),
			wantPanic: false,
		},
		{
			name: "парсинг строки с разными цифрами",
			args: args{
				s: "94321",
			},
			want:      []int{9, 4, 3, 2, 1},
			wantPanic: false,
		},
		{
			name: "парсинг пустой строки",
			args: args{
				s: "",
			},
			want:      []int{},
			wantPanic: false,
		},
		{
			name: "парсинг строки с нечисловым значением",
			args: args{
				s: "abcd123",
			},
			want:      []int{},
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

			if got := parseDigits(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
