package inn

import "testing"

func Test_validatePersonINN(t *testing.T) {
	type args struct {
		inn string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
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
			if err := validatePersonINN(tt.args.inn); (err != nil) != tt.wantErr {
				t.Errorf("validatePersonINN() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
