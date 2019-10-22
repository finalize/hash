package interfaces

import (
	"database/sql"

	"github.com/shgysd/hash/api/model"
	"github.com/shgysd/hash/api/repository"
)

// NewUserRepo Initialize user repository
func NewUserRepo(conn *sql.DB) repository.UserRepository {
	return &UserRepository{
		Conn: conn,
	}
}

// UserRepository Handler with DB
type UserRepository struct {
	Conn *sql.DB
}

// SignUp SignUp
func (h *UserRepository) SignUp(u *model.User) {

	// Validate
	// if u.Password == "" {
	// 	return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid cred")
	// }

	// pwd := []byte(u.Password)
	// hash := common.HashAndSalt(pwd)
	// hashID := u.HashID

	// u.HashID = hashID
	// u.Password = hash
	// if !h.Conn.NewRecord(&u) {
	// 	panic("could not create new record")
	// }
	// if err := h.Conn.Create(&u).Error; err != nil {
	// 	panic(err.Error())
	// }

}
