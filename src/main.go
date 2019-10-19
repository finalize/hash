package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func handler(w http.ResponseWriter, r *http.Request) {
	res := map[string]string{"status": "ok"}

	json, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func main() {
	connectionString := getConnectionString()
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	http.HandleFunc("/", handler)
	fmt.Println("server is listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

func getConnectionString() string {
	host := getParamString("MYSQL_DB_HOST", "localhost")
	port := getParamString("MYSQL_PORT", "3306")
	user := getParamString("MYSQL_USER", "root")
	pass := getParamString("MYSQL_PASSWORD", "mysql")
	dbname := getParamString("MYSQL_DB", "test_db")
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
