package mysql

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/albinzx/sql/tls"
	"github.com/go-sql-driver/mysql"
)

const (
	// DriverMySQL is driver name for mysql
	DriverMySQL = "mysql"
)

// DataSource is mysql data source
type DataSource struct {
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

// dsn returns mysql data source name
func (my *DataSource) dsn(tlsKey string) string {
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", my.User, my.Password, my.Host, my.Port, my.Database)
	val := url.Values{}

	if my.ParseTime {
		val.Add("parseTime", "1")
	}
	if len(my.Location) > 0 {
		val.Add("loc", my.Location)
	}
	if my.Timeout > 0 {
		val.Add("timeout", my.Timeout.String())
	}
	if len(tlsKey) > 0 {
		val.Add("tls", tlsKey)
	}

	if len(val) == 0 {
		return connection
	}
	return fmt.Sprintf("%s?%s", connection, val.Encode())
}

// Name returns mysql driver name and data source name
func (my *DataSource) Name() (string, string, error) {

	var tlsKey string
	if len(my.CA) > 0 && my.ServerName != "" {
		tlsKey = "custom"
		if err := mysql.RegisterTLSConfig(tlsKey, tls.WithServerAndCA(my.ServerName, my.CA)); err != nil {
			log.Printf("error while registering tls config: %v", err)

			return "", "", err
		}
	} else if len(my.CA) > 0 {
		tlsKey = "custom"
		if err := mysql.RegisterTLSConfig(tlsKey, tls.WithCA(my.CA)); err != nil {
			log.Printf("error while registering tls config: %v", err)

			return "", "", err
		}
	}

	return DriverMySQL, my.dsn(tlsKey), nil

}

// Driver returns mysql driver name
func (my *DataSource) Driver() string {
	return DriverMySQL
}
