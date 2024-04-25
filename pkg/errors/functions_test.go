package errors

import (
	"errors"
	"testing"
)

func TestErrJoin(t *testing.T) {
	type args struct {
		errs []error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy path",
			args: args{
				errs: []error{
					errors.New("error 1"),
					errors.New("error 2"),
					errors.Join(errors.New("error 3"), errors.New("error 4")),
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ErrJoin(tt.args.errs...); (err != nil) != tt.wantErr {
				t.Errorf("ErrJoin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
