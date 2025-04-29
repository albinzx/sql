package mysql

import (
	"encoding/base64"
	"testing"
	"time"
)

func ca() []byte {
	strCA := "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUN1RENDQWwrZ0F3SUJBZ0lVR2hhdkpudzZTV2RJRFp2bVpQSWJyUjFETzJvd0NnWUlLb1pJemowRUF3SXcKZ2JFeEN6QUpCZ05WQkFZVEFrbEVNUlF3RWdZRFZRUUlEQXRFUzBrZ1NtRnJZWEowWVRFWU1CWUdBMVVFQnd3UApTbUZyWVhKMFlTQlRaV3hoZEdGdU1TRXdId1lEVlFRS0RCaFFWQ0JEWVhCcGRHRnNJRTVsZENCSmJtUnZibVZ6CmFXRXhFekFSQmdOVkJBc01DbFJsWTJodWIyeHZaM2t4RlRBVEJnTlZCQU1NREdSbGRpNWthVzFwYVM1cFpERWoKTUNFR0NTcUdTSWIzRFFFSkFSWVVjM2x6WVdSdGFXNUFZMkZ3YVhSaGJIZ3VhV1F3SGhjTk1qQXdOakE1TVRJMApOVEV3V2hjTk16QXdOakEzTVRJME5URXdXakNCc1RFTE1Ba0dBMVVFQmhNQ1NVUXhGREFTQmdOVkJBZ01DMFJMClNTQktZV3RoY25SaE1SZ3dGZ1lEVlFRSERBOUtZV3RoY25SaElGTmxiR0YwWVc0eElUQWZCZ05WQkFvTUdGQlUKSUVOaGNHbDBZV3dnVG1WMElFbHVaRzl1WlhOcFlURVRNQkVHQTFVRUN3d0tWR1ZqYUc1dmJHOW5lVEVWTUJNRwpBMVVFQXd3TVpHVjJMbVJwYldscExtbGtNU013SVFZSktvWklodmNOQVFrQkZoUnplWE5oWkcxcGJrQmpZWEJwCmRHRnNlQzVwWkRCWk1CTUdCeXFHU000OUFnRUdDQ3FHU000OUF3RUhBMElBQkdwZnZyL2doakVVK25uakRHakYKVXp1WUZSUE1JRzF5M0lXcVozMytYR0tTbERpTnVlclBFZGFuaG5ZclJWazd3b3RoNTZhYWp5VDVPbE9nakJTawpVWFdqVXpCUk1CMEdBMVVkRGdRV0JCUVBwS1F4bUl6N1dOOVZCVXVRbkxZMENnS2ppekFmQmdOVkhTTUVHREFXCmdCUVBwS1F4bUl6N1dOOVZCVXVRbkxZMENnS2ppekFQQmdOVkhSTUJBZjhFQlRBREFRSC9NQW9HQ0NxR1NNNDkKQkFNQ0EwY0FNRVFDSUhzMVFkVHl2U05PUE9NSWlMWERFZFA3dWl6VitEc0lwOGpzNStyMXVoMVZBaUJnMnRzcwpJYjJ2UGd3T1BJaVhXakVGZmxrdkFvakJ1WVJEdmNKbDBFcXR2Zz09Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"
	ca, _ := base64.StdEncoding.DecodeString(strCA)

	return ca
}

func TestDataSource_Name(t *testing.T) {
	type fields struct {
		Host       string
		Port       string
		User       string
		Password   string
		Database   string
		CA         []byte
		ServerName string
		ParseTime  bool
		Location   string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		want1   string
		wantErr bool
	}{
		{name: "mysql", fields: fields{
			Host:       "localhost",
			Port:       "3306",
			User:       "user",
			Password:   "pass",
			Database:   "db",
			CA:         ca(),
			ServerName: "mysql.db.id",
			ParseTime:  true,
			Location:   "Asia/Jakarta",
		}, want: "mysql",
			want1:   "user:pass@tcp(localhost:3306)/db?loc=Asia%2FJakarta&parseTime=1&tls=custom",
			wantErr: false},
		{name: "mysql with empty server name", fields: fields{
			Host:      "localhost",
			Port:      "3306",
			User:      "user",
			Password:  "pass",
			Database:  "db",
			CA:        ca(),
			ParseTime: true,
			Location:  "Asia/Jakarta",
		}, want: "mysql",
			want1:   "user:pass@tcp(localhost:3306)/db?loc=Asia%2FJakarta&parseTime=1&tls=custom",
			wantErr: false},
		{name: "mysql with empty ca", fields: fields{
			Host:      "localhost",
			Port:      "3306",
			User:      "user",
			Password:  "pass",
			Database:  "db",
			CA:        []byte{},
			ParseTime: true,
			Location:  "Asia/Jakarta",
		}, want: "mysql",
			want1:   "user:pass@tcp(localhost:3306)/db?loc=Asia%2FJakarta&parseTime=1",
			wantErr: false},
		{name: "mysql with nil ca", fields: fields{
			Host:     "localhost",
			Port:     "3306",
			User:     "user",
			Password: "pass",
			Database: "db",
		}, want: "mysql",
			want1:   "user:pass@tcp(localhost:3306)/db",
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			my := &DataSource{
				Host:       tt.fields.Host,
				Port:       tt.fields.Port,
				User:       tt.fields.User,
				Password:   tt.fields.Password,
				Database:   tt.fields.Database,
				CA:         tt.fields.CA,
				ServerName: tt.fields.ServerName,
				ParseTime:  tt.fields.ParseTime,
				Location:   tt.fields.Location,
			}
			got, got1, err := my.Name()
			if (err != nil) != tt.wantErr {
				t.Errorf("Name() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Name() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Name() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDataSource_dsn(t *testing.T) {
	type fields struct {
		Host       string
		Port       string
		User       string
		Password   string
		Database   string
		CA         []byte
		ServerName string
		ParseTime  bool
		Location   string
		Timeout    time.Duration
	}
	type args struct {
		tlsKey string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "full dsn",
			fields: fields{
				Host:      "localhost",
				Port:      "3306",
				User:      "user",
				Password:  "pass",
				Database:  "db",
				ParseTime: true,
				Location:  "Asia/Jakarta",
			},
			args: args{
				tlsKey: "custom",
			},
			want: "user:pass@tcp(localhost:3306)/db?loc=Asia%2FJakarta&parseTime=1&tls=custom",
		},
		{
			name: "simple dsn",
			fields: fields{
				Host:     "localhost",
				Port:     "3306",
				User:     "user",
				Password: "pass",
				Database: "db",
			},
			want: "user:pass@tcp(localhost:3306)/db",
		},
		{
			name: "full dsn with timeout",
			fields: fields{
				Host:      "localhost",
				Port:      "3306",
				User:      "user",
				Password:  "pass",
				Database:  "db",
				ParseTime: true,
				Location:  "Asia/Jakarta",
				Timeout:   3 * time.Second,
			},
			args: args{
				tlsKey: "custom",
			},
			want: "user:pass@tcp(localhost:3306)/db?loc=Asia%2FJakarta&parseTime=1&timeout=3s&tls=custom",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			my := &DataSource{
				Host:       tt.fields.Host,
				Port:       tt.fields.Port,
				User:       tt.fields.User,
				Password:   tt.fields.Password,
				Database:   tt.fields.Database,
				CA:         tt.fields.CA,
				ServerName: tt.fields.ServerName,
				ParseTime:  tt.fields.ParseTime,
				Location:   tt.fields.Location,
				Timeout:    tt.fields.Timeout,
			}
			if got := my.dsn(tt.args.tlsKey); got != tt.want {
				t.Errorf("dsn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDataSource_Driver(t *testing.T) {
	type fields struct {
		Host       string
		Port       string
		User       string
		Password   string
		Database   string
		CA         []byte
		ServerName string
		ParseTime  bool
		Location   string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "mysql driver", fields: fields{}, want: DriverMySQL},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			my := &DataSource{
				Host:       tt.fields.Host,
				Port:       tt.fields.Port,
				User:       tt.fields.User,
				Password:   tt.fields.Password,
				Database:   tt.fields.Database,
				CA:         tt.fields.CA,
				ServerName: tt.fields.ServerName,
				ParseTime:  tt.fields.ParseTime,
				Location:   tt.fields.Location,
			}
			if got := my.Driver(); got != tt.want {
				t.Errorf("Driver() = %v, want %v", got, tt.want)
			}
		})
	}
}
