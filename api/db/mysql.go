package db

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/shgysd/hash/api/config"
)

var db *sql.DB

// New returns sql.DB
func New(d *config.Config) *sql.DB {
	connectionString := getConnectionString(d)
	conn, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error())
	}

	err = conn.Ping()
	if err != nil {
		panic(err.Error())
	}

	db = conn

	return db
}

// GetDB returns sql.DB
func GetDB() *sql.DB {
	return db
}

func getConnectionString(d *config.Config) string {
	host := getParamString("MYSQL_DB_HOST", "mysql")
	port := getParamString("MYSQL_PORT", "3306")
	user := getParamString("MYSQL_USER", d.MySQL.User)
	pass := getParamString("MYSQL_PASSWORD", d.MySQL.Password)
	dbname := getParamString("MYSQL_DB", d.MySQL.Name)
	protocol := getParamString("MYSQL_PROTOCOL", "tcp")
	dbargs := getParamString("MYSQL_DBARGS", " ")

	if strings.Trim(dbargs, " ") != "" {
		dbargs = "?" + dbargs
	} else {
		dbargs = ""
	}
	return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s",
		user, pass, protocol, host, port, dbname, dbargs)
}

func getParamString(param string, defaultValue string) string {
	env := os.Getenv(param)
	if env != "" {
		return env
	}
	return defaultValue
}
