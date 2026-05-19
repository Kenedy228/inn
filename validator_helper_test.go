package inn

import (
	"strings"
	"testing"
)

func Test_validateDigitsOnly(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "только цифры",
			args: args{
				code: "123456",
			},
			wantErr: false,
		},
		{
			name: "цифры + пробелы",
			args: args{
				code: "123 345",
			},
			wantErr: true,
		},
		{
			name: "цифры + буквы",
			args: args{
				code: "123dsa",
			},
			wantErr: true,
		},
		{
			name: "пустое значение",
			args: args{
				code: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateDigitsOnly(tt.args.code); (err != nil) != tt.wantErr {
				t.Errorf("validateDigitsOnly() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validateCodeLength(t *testing.T) {
	type args struct {
		code           string
		expectedLength int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "разная длина",
			args: args{
				code:           "абвгд",
				expectedLength: 3,
			},
			wantErr: true,
		},
		{
			name: "одинаковая длина",
			args: args{
				code:           "абвгд",
				expectedLength: 5,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateCodeLength(tt.args.code, tt.args.expectedLength); (err != nil) != tt.wantErr {
				t.Errorf("validateCodeLength() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validateChecksumDigit(t *testing.T) {
	type args struct {
		baseDigits    []int
		weights       []int
		expectedDigit int
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ФЛ/ИП: первая контрольная цифра валидна",
			args: args{
				// ИНН: 500100732259
				baseDigits:    []int{5, 0, 0, 1, 0, 0, 7, 3, 2, 2},
				weights:       []int{7, 2, 4, 10, 3, 5, 9, 4, 6, 8},
				expectedDigit: 5,
			},
			wantErr: false,
		},
		{
			name: "ФЛ/ИП: вторая контрольная цифра валидна",
			args: args{
				// ИНН: 500100732259
				baseDigits:    []int{5, 0, 0, 1, 0, 0, 7, 3, 2, 2, 5},
				weights:       []int{3, 7, 2, 4, 10, 3, 5, 9, 4, 6, 8},
				expectedDigit: 9,
			},
			wantErr: false,
		},
		{
			name: "ЮЛ: контрольная цифра валидна",
			args: args{
				// ИНН: 7707083893
				baseDigits:    []int{7, 7, 0, 7, 0, 8, 3, 8, 9},
				weights:       []int{2, 4, 10, 3, 5, 9, 4, 6, 8},
				expectedDigit: 3,
			},
			wantErr: false,
		},
		{
			name: "ФЛ/ИП: первая контрольная цифра невалидна",
			args: args{
				// ИНН: 500100732259, но ожидаемая цифра испорчена
				baseDigits:    []int{5, 0, 0, 1, 0, 0, 7, 3, 2, 2},
				weights:       []int{7, 2, 4, 10, 3, 5, 9, 4, 6, 8},
				expectedDigit: 0,
			},
			wantErr: true,
		},
		{
			name: "ФЛ/ИП: вторая контрольная цифра невалидна",
			args: args{
				// ИНН: 500100732259, но ожидаемая цифра испорчена
				baseDigits:    []int{5, 0, 0, 1, 0, 0, 7, 3, 2, 2, 5},
				weights:       []int{3, 7, 2, 4, 10, 3, 5, 9, 4, 6, 8},
				expectedDigit: 0,
			},
			wantErr: true,
		},
		{
			name: "ЮЛ: контрольная цифра невалидна",
			args: args{
				// ИНН: 7707083893, но ожидаемая цифра испорчена
				baseDigits:    []int{7, 7, 0, 7, 0, 8, 3, 8, 9},
				weights:       []int{2, 4, 10, 3, 5, 9, 4, 6, 8},
				expectedDigit: 0,
			},
			wantErr: true,
		},
		{
			name: "остаток 10 превращается в 0",
			args: args{
				// weightedSum = 21, 21 % 11 = 10, 10 % 10 = 0
				baseDigits:    []int{3},
				weights:       []int{7},
				expectedDigit: 0,
			},
			wantErr: false,
		},
		{
			name: "остаток 10 превращается в 0, но ожидали другую цифру",
			args: args{
				baseDigits:    []int{3},
				weights:       []int{7},
				expectedDigit: 1,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateChecksumDigit(tt.args.baseDigits, tt.args.weights, tt.args.expectedDigit); (err != nil) != tt.wantErr {
				t.Errorf("validateChecksumDigit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validateForbiddenValue(t *testing.T) {
	type args struct {
		inn       string
		forbidden string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "запрещенное значение",
			args: args{
				inn:       strings.Repeat("0", 10),
				forbidden: strings.Repeat("0", 10),
			},
			wantErr: true,
		},
		{
			name: "разная длина ИНН и запрещенного значения",
			args: args{
				inn:       strings.Repeat("0", 9),
				forbidden: strings.Repeat("0", 10),
			},
			wantErr: false,
		},
		{
			name: "незапрещенное значение",
			args: args{
				inn:       strings.Repeat("1", 10),
				forbidden: strings.Repeat("0", 10),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateForbiddenValue(tt.args.inn, tt.args.forbidden); (err != nil) != tt.wantErr {
				t.Errorf("validateForbiddenValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
