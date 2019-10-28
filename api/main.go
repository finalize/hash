package main

import (
	"fmt"
	"net/http"

	"github.com/shgysd/hash/api/config"
	"github.com/shgysd/hash/api/db"
	"github.com/shgysd/hash/api/handler"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	config := config.New()

	db.New(config)
}

func main() {
	db := db.GetDB()
	defer db.Close()

	mux := http.NewServeMux()
	handler.NewRouter(mux, db)
	fmt.Println("server is listening on port 8080")
	http.ListenAndServe(":8080", mux)
}
