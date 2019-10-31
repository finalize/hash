package main

import (
	"github.com/labstack/echo"
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

	e := echo.New()
	handler.NewRouter(e, db)
	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
