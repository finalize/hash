package handler

import (
	"database/sql"

	"github.com/labstack/echo"
)

// NewRouter Init Router
func NewRouter(e *echo.Echo, db *sql.DB) *echo.Echo {
	// config := config.New()
	user := NewUserHandler(db)
	// tag := NewTagHandler(db)
	// auth := NewOAuthHandler(config, db)

	// Auth
	e.POST("/signup", user.SignUp)
	e.POST("/signin", user.SignIn)
	// e.POST("/login", user.Login)
	// // Social login
	// e.GET("/auth/twitter", auth.TwitterLogin())
	// e.GET("/auth/twitter/callback", auth.TwitterCallback())
	// // User
	// e.GET("/users/:id", user.GetUser)
	// e.POST("/users/:id/tags", user.CreateTag)
	// // Tag
	// e.GET("/tags/:name/users", tag.GetUsers)
	return e
}
