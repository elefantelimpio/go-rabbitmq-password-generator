package pwdgenerator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPwdService(t *testing.T) {
	tests := []struct {
		name string
		want PwdIService
	}{
		{
			name: "all ok",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotPanics(t, func() { NewPwdService() })
		})
	}
}

func Test_pwdService_GetSHA256Hash(t *testing.T) {
	type args struct {
		pwd string
	}
	tests := []struct {
		name    string
		srv     pwdService
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "all ok",
			args: args{
				pwd: "testpwd",
			},
			srv:     pwdService{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := pwdService{}
			_, err := srv.GetSHA256Hash(tt.args.pwd)
			if (err != nil) != tt.wantErr {
				t.Errorf("pwdService.GetSHA256Hash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_getPwdSHA256Hash(t *testing.T) {
	type args struct {
		salt string
		pwd  string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "all ok",
			args: args{
				salt: "AB12",
				pwd:  "AB12",
			},
			want:    "e88371a11c0760eb0da7da45b5b8d41698d0084a4218d9061972a43484c00859",
			wantErr: false,
		},
		{
			name: "all ok",
			args: args{
				salt: "123",
				pwd:  "AS",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getPwdSHA256Hash(tt.args.salt, tt.args.pwd)
			if (err != nil) != tt.wantErr {
				t.Errorf("getPwdSHA256Hash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getPwdSHA256Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_encodeToBase64(t *testing.T) {
	type args struct {
		toEncode string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "all ok",
			args: args{
				toEncode: "ADBC",
			},
			want:    "rbw",
			wantErr: false,
		},
		{
			name: "error",
			args: args{
				toEncode: "asd",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := encodeToBase64(tt.args.toEncode)
			if (err != nil) != tt.wantErr {
				t.Errorf("encodeToBase64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("encodeToBase64() = %v, want %v", got, tt.want)
			}
		})
	}
}
