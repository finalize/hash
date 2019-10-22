package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type apiHandler struct{}

type (
	// Handler Handler
	Handler struct {
		DB *sql.DB
	}
)

const (
	// Key This should be imported from somewhere else
	Key = "secret"
)

// InitializeRouter Init Router
func InitializeRouter(db *sql.DB, mux *http.ServeMux) *http.ServeMux {
	user := NewUserHandler(db)

	mux.HandleFunc("/signup", user.SignUp)
	return mux
}

func handler(w http.ResponseWriter, req *http.Request) {

	resp := map[string]int{"user_id": 1}

	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(resp)
}
