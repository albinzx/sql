package ha

import (
	"database/sql"
	"reflect"
	"testing"
)

func TestDB(t *testing.T) {
	type args struct {
		primaries []*sql.DB
		replicas  []*sql.DB
	}
	dbp, _ := sql.Open("mysql", "test:test@tcp(localhost:3306)/primary")
	dbr, _ := sql.Open("mysql", "test:test@tcp(localhost:3306)/replica")
	tests := []struct {
		name    string
		args    args
		wantNil bool
	}{
		{name: "new", args: args{
			primaries: []*sql.DB{dbp},
			replicas:  []*sql.DB{dbr},
		}, wantNil: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DB(tt.args.primaries, tt.args.replicas); !reflect.DeepEqual(got == nil, tt.wantNil) {
				t.Errorf("DB() = %v, want nil %v", got, tt.wantNil)
			}
		})
	}
}
