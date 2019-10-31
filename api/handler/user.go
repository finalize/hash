package handler

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/shgysd/hash/api/interfaces"
	"github.com/shgysd/hash/api/repository"
	"github.com/shgysd/hash/api/utils/auth"

	"github.com/labstack/echo"
)

// UserHandler contains user interface
type UserHandler struct {
	repo repository.UserRepository
}

type signUp struct {
	Name        string `json:"name" validate:"required"`
	DisplayName string `json:"displayName" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required"`
}

// NewUserHandler inits user handler
func NewUserHandler(conn *sql.DB) *UserHandler {
	return &UserHandler{
		repo: interfaces.NewUserRepo(conn),
	}
}

// SignUp creates a new user
func (h *UserHandler) SignUp(c echo.Context) (err error) {
	data := &repository.SignUp{}
	if err := c.Bind(data); err != nil {
		return err
	}

	id := h.repo.SignUp(data)

	tokenString, err := auth.CreateJSONWebToken(id)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	resp := map[string]interface{}{"token": tokenString}
	return c.JSON(http.StatusCreated, resp)
}
