package db

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/shgysd/hash/api/config"
)

var db *sql.DB

// New Connect to MySQL
func New(d *config.Config) *sql.DB {
	connectionString := getConnectionString(d)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error())
	}

	return db
}

// Init Migration
func Init() {
	return
}

// GetDB for getting db
func GetDB() *sql.DB {
	return db
}

func getConnectionString(d *config.Config) string {
	host := getParamString("MYSQL_DB_HOST", d.MySQL.Host)
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
