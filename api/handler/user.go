package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/shgysd/hash/api/interfaces"
	"github.com/shgysd/hash/api/model"
	"github.com/shgysd/hash/api/repository"
)

// NewUserHandler Initialize user repository
func NewUserHandler(conn *sql.DB) *UserHandler {
	return &UserHandler{
		repo: interfaces.NewUserRepo(conn),
	}
}

// UserHandler Handler with DB
type UserHandler struct {
	repo repository.UserRepository
}

type Message struct {
	Data interface{} `json:"data"`
}

// SignUp sign up
func (h *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var msg interface{}
	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Println(msg)

	var u model.User
	h.repo.SignUp(&u)

	resp := map[string]int{"user_id": 2}
	// if err := h.repo.SignUp(u); err != nil {
	// 	return err
	// }

	// if u.Email != "" || u.Password != "" {

	// 	token := jwt.New(jwt.SigningMethodHS256)
	// 	claims := token.Claims.(jwt.MapClaims)
	// 	claims["admin"] = true
	// 	claims["hashID"] = u.HashID
	// 	claims["displayName"] = u.DisplayName
	// 	claims["iat"] = time.Now()
	// 	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	// 	tokenString, _ := token.SignedString([]byte(Key))
	// 	data := map[string]interface{}{"token": tokenString}

	// 	return c.JSON(http.StatusCreated, data)
	// }

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(resp)
}