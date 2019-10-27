package handler

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/shgysd/hash/api/interfaces"
	"github.com/shgysd/hash/api/repository"
	"github.com/shgysd/hash/api/types"
	"gopkg.in/go-playground/validator.v9"
)

// UserHandler Handler with DB
type UserHandler struct {
	repo repository.UserRepository
}

// NewUserHandler Initialize user repository
func NewUserHandler(conn *sql.DB) *UserHandler {
	return &UserHandler{
		repo: interfaces.NewUserRepo(conn),
	}
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
	var data types.SignUp
	err = json.Unmarshal(b, &data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	validate := validator.New() //インスタンス生成
	errors := validate.Struct(data)
	if errors != nil {
		log.Println(errors.Error())
	}

	id := h.repo.SignUp(&data)

	resp := map[string]interface{}{"user_id": id}
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
