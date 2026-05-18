package inn

import "testing"

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
