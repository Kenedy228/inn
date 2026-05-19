package inn

import (
	"strings"
	"testing"
)

func Test_validateLegalEntityINN(t *testing.T) {
	type args struct {
		inn string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "валидный ИНН",
			args: args{
				inn: "1030000000",
			},
			wantErr: false,
		},
		{
			name: "невалидный ИНН",
			args: args{
				inn: "7728168972", // Правильная 1, заменили на 2
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateLegalEntityINN(tt.args.inn); (err != nil) != tt.wantErr {
				t.Errorf("validateLegalEntityINN() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateNaturalPerson(t *testing.T) {
	type args struct {
		inn string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "пустая строка",
			args: args{
				inn: "",
			},
			wantErr: true,
		},
		{
			name: "только пробелы",
			args: args{
				inn: strings.Repeat(" ", 10),
			},
			wantErr: true,
		},
		{
			name: "содержит нецифры",
			args: args{
				inn: "123adsb-cds",
			},
			wantErr: true,
		},
		{
			name: "длина меньше нужной",
			args: args{
				inn: strings.Repeat("1", naturalPersonDigitsCount-1),
			},
			wantErr: true,
		},
		{
			name: "длина больше нужной",
			args: args{
				inn: strings.Repeat("1", naturalPersonDigitsCount+1),
			},
			wantErr: true,
		},
		{
			name: "запрещенное значение",
			args: args{
				inn: strings.Repeat("0", naturalPersonDigitsCount),
			},
			wantErr: true,
		},
		{
			name: "валидный ИНН для ФЛ",
			args: args{
				inn: "500100732259",
			},
			wantErr: false,
		},
		{
			name: "неверная вторая контрольная цифра",
			args: args{
				inn: "500100732249", // правильная 5, заменили на 4
			},
			wantErr: true,
		},
		{
			name: "неверная вторая контрольная цифра",
			args: args{
				inn: "500100732258", // правильная 9, заменили на 8
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateNaturalPerson(tt.args.inn); (err != nil) != tt.wantErr {
				t.Errorf("ValidateIndividualEntrepreneur() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateIndividualEntrepreneur(t *testing.T) {
	type args struct {
		inn string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "пустая строка",
			args: args{
				inn: "",
			},
			wantErr: true,
		},
		{
			name: "только пробелы",
			args: args{
				inn: strings.Repeat(" ", 10),
			},
			wantErr: true,
		},
		{
			name: "содержит нецифры",
			args: args{
				inn: "123adsb-cds",
			},
			wantErr: true,
		},
		{
			name: "длина меньше нужной",
			args: args{
				inn: strings.Repeat("1", individualEntrepreneurDigitsCount-1),
			},
			wantErr: true,
		},
		{
			name: "длина больше нужной",
			args: args{
				inn: strings.Repeat("1", individualEntrepreneurDigitsCount+1),
			},
			wantErr: true,
		},
		{
			name: "запрещенное значение",
			args: args{
				inn: strings.Repeat("0", individualEntrepreneurDigitsCount),
			},
			wantErr: true,
		},
		{
			name: "валидный ИНН для ИП",
			args: args{
				inn: "500100732259",
			},
			wantErr: false,
		},
		{
			name: "неверная вторая контрольная цифра",
			args: args{
				inn: "500100732249", // правильная 5, заменили на 4
			},
			wantErr: true,
		},
		{
			name: "неверная вторая контрольная цифра",
			args: args{
				inn: "500100732258", // правильная 9, заменили на 8
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateIndividualEntrepreneur(tt.args.inn); (err != nil) != tt.wantErr {
				t.Errorf("ValidateIndividualEntrepreneur() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
