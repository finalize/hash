package handler

import (
	"database/sql"
	"net/http"
)

// NewRouter Init Router
func NewRouter(mux *http.ServeMux, db *sql.DB) *http.ServeMux {
	user := NewUserHandler(db)

	mux.Handle("/signup", JWTMiddleware.Handler(http.HandlerFunc(user.SignUp)))
	mux.Handle("/signin", http.HandlerFunc(user.SignIn))
	mux.Handle("/users/", http.HandlerFunc(user.GetUser))
	mux.Handle("/token", GetTokenHandler)
	return mux
}
