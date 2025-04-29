package sql

import (
	"errors"
	"reflect"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type dummyDataSource struct{}

func (d *dummyDataSource) Driver() string {
	return "mysql"
}

func (d *dummyDataSource) Name() (string, string, error) {
	return "mysql", "test:test@tcp(localhost:3306)/test", nil
}

type errorDataSourceName struct{}

func (e *errorDataSourceName) Driver() string {
	return ""
}

func (e *errorDataSourceName) Name() (string, string, error) {
	return "", "", errors.New("data source error")
}

type errorDataSource struct{}

func (d *errorDataSource) Driver() string {
	return "unknown"
}

func (d *errorDataSource) Name() (string, string, error) {
	return "unknown", "test:test@tcp(localhost:3306)/test", nil
}

func TestDB(t *testing.T) {
	type args struct {
		source  DataSource
		options []Option
	}
	tests := []struct {
		name    string
		args    args
		wantNil bool
		wantErr bool
	}{
		{name: "db ok", args: args{
			source:  &dummyDataSource{},
			options: []Option{WithConnection(5, 5, 5*time.Minute, 3*time.Minute)},
		}, wantNil: false, wantErr: false},
		{name: "db name error", args: args{
			source:  &errorDataSourceName{},
			options: []Option{WithConnection(5, 5, 5*time.Minute, 3*time.Minute)},
		}, wantNil: true, wantErr: true},
		{name: "db error", args: args{
			source:  &errorDataSource{},
			options: []Option{WithConnection(5, 5, 5*time.Minute, 3*time.Minute)},
		}, wantNil: true, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DB(tt.args.source, tt.args.options...)
			if (err != nil) != tt.wantErr {
				t.Errorf("DB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (got == nil) != tt.wantNil {
				t.Errorf("DB() got = %v, wantNil %v", got, tt.wantNil)
			}
		})
	}
}

func TestWithConnection(t *testing.T) {
	type args struct {
		maxOpen     int
		maxIdle     int
		maxLifetime time.Duration
		maxIdleTime time.Duration
	}
	tests := []struct {
		name string
		args args
		want *config
	}{
		{name: "with connection", args: args{
			maxOpen:     5,
			maxIdle:     5,
			maxLifetime: 5 * time.Hour,
			maxIdleTime: 5 * time.Minute,
		}, want: &config{
			maxOpen:     5,
			maxIdle:     5,
			maxLifetime: 5 * time.Hour,
			maxIdleTime: 5 * time.Minute,
		}},
		{name: "with connection max idle exceed max open", args: args{
			maxOpen:     5,
			maxIdle:     10,
			maxLifetime: 5 * time.Hour,
			maxIdleTime: 5 * time.Minute,
		}, want: &config{
			maxOpen:     5,
			maxIdle:     5,
			maxLifetime: 5 * time.Hour,
			maxIdleTime: 5 * time.Minute,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &config{}
			WithConnection(tt.args.maxOpen, tt.args.maxIdle, tt.args.maxLifetime, tt.args.maxIdleTime)(cfg)
			if got := cfg; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithConnection() got %+v, want %+v", got, tt.want)
			}
		})
	}
}

func Test_defaults(t *testing.T) {
	type args struct {
		config *config
	}
	tests := []struct {
		name string
		args args
		want *config
	}{
		{name: "defaults", args: args{
			config: &config{}},
			want: &config{
				maxOpen:     defaultMaxOpen,
				maxIdle:     defaultMaxIdle,
				maxLifetime: defaultMaxLifetime,
				maxIdleTime: defaultMaxIdleTime,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defaults(tt.args.config)
			if got := tt.args.config; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("defaults() got %+v, want %+v", got, tt.want)
			}
		})
	}
}
