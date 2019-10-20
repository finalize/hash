package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
)

type apiHandler struct{}

type (
	// Handler Handler
	Handler struct {
		DB *gorm.DB
	}
)

const (
	// Key This should be imported from somewhere else
	Key = "secret"
)

// InitializeRouter Init Router
func InitializeRouter(db *sql.DB, mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc("/", handler)
	return mux
}

func handler(w http.ResponseWriter, req *http.Request) {

	fmt.Println(req)

	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}
	fmt.Fprintf(w, "Welcome to the home page!")
}
