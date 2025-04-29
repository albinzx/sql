# SQL

SQL library for sql db.

## Usage

```go
package main

import (
	"time"

	"github.com/albinzx/sql"
	"github.com/albinzx/sql/mysql"
)

func main() {
	ds := &mysql.DataSource{
		Host:     "localhost",
		Port:     "3306",
		User:     "user",
		Password: "pass",
		Database: "db",
	}

	db, _ := sql.DB(ds,
		sql.WithConnection(10, 10, time.Hour, 10*time.Minute))

	db.Ping()
}
```
