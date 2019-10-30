package handler

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/shgysd/hash/api/utils"
	"github.com/shgysd/hash/api/utils/auth"
	"github.com/shgysd/hash/api/interfaces"
	"github.com/shgysd/hash/api/repository"
	"github.com/shgysd/hash/api/types"
	"gopkg.in/go-playground/validator.v9"
)

// UserHandler user handler
type UserHandler struct {
	repo repository.UserRepository
}

// NewUserHandler Initialize user handler
func NewUserHandler(conn *sql.DB) *UserHandler {
	return &UserHandler{
		repo: interfaces.NewUserRepo(conn),
	}
}

// SignUp crate a user
func (h *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	if utils.IsAllowMethod(w, r.Method, "POST") {
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var data types.SignUp
	err = json.Unmarshal(b, &data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	validate := validator.New()
	err = validate.Struct(data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	// Insert data to db
	id := h.repo.SignUp(&data)

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["admin"] = true
	claims["sub"] = id
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Create signed token
	tokenString, err := token.SignedString([]byte(os.Getenv("KEY")))
	if err != nil {
		log.Println(err.Error())
	}

	resp := map[string]interface{}{"token": tokenString}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(resp)
	return
}

// SignIn login
func (h *UserHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	if utils.IsAllowMethod(w, r.Method, "POST") {
		return
	}

	var data types.SignIn

	utils.UnmarshalBody(r.Body, &data)

	validate := validator.New()
	err := validate.Struct(data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	// Insert data to db
	id := h.repo.SignIn(&data)

	// Create signed token
	tokenString, err := auth.CreateJSONWebToken(id)
	if err != nil {
		log.Println(err.Error())
	}
	resp := map[string]interface{}{"token": tokenString}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(resp)
	return
}

// GetUser get a user
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {

}
