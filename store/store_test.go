package store

import "testing"

func TestIsCorrectName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "1",
			args: args{name: "Felix Peterson"},
			wantErr: false,
		},
		{
			name: "2",
			args: args{name: "Felix Van Peterson"},
			wantErr: false,
		},
		{
			name: "3",
			args: args{name: "Felix -Van Peterson"},
			wantErr: true,
		},
		{
			name: "4",
			args: args{name: "Felix"},
			wantErr: true,
		},
		{
			name: "5",
			args: args{name: "Felix @me"},
			wantErr: true,
		},
		{
			name: "6",
			args: args{name: ""},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := IsCorrectName(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("IsCorrectName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
