package inn

import "testing"

func TestType_isValid(t *testing.T) {
	tests := []struct {
		name string
		t    Type
		want bool
	}{
		{
			name: "невалидный тип",
			t:    Type("invalid"),
			want: false,
		},
		{
			name: "пустой тип",
			t:    Type(""),
			want: false,
		},
		{
			name: "пустая строка",
			t:    "",
			want: false,
		},
		{
			name: "валидный тип",
			t:    TypeIP,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.isValid(); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestType_codeLength(t *testing.T) {
	tests := []struct {
		name string
		t    Type
		want int
	}{
		{
			name: "валидное значение",
			t:    TypeIP,
			want: 12,
		},
		{
			name: "невалидное значение",
			t:    "",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.codeLength(); got != tt.want {
				t.Errorf("codeLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
