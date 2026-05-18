package inn

import "testing"

func TestNormalize(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "удаляет все пробелы",
			args: args{
				code: " 11 22    33 44      ",
			},
			want: "11223344",
		},
		{
			name: "не изменяет значение без пробелов",
			args: args{
				code: "11223344",
			},
			want: "11223344",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Normalize(tt.args.code); got != tt.want {
				t.Errorf("Normalize() = %v, want %v", got, tt.want)
			}
		})
	}
}
