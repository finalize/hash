package interfaces

import (
	"log"

	"database/sql"

	"github.com/shgysd/hash/api/repository"
	"github.com/shgysd/hash/api/utils/crypto"
)

// UserRepository contains db
type UserRepository struct {
	Conn *sql.DB
}

type signUp struct {
	Name        string `json:"name" validate:"required"`
	DisplayName string `json:"displayName" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required"`
}

// NewUserRepo returns user repository that contains db
func NewUserRepo(conn *sql.DB) repository.UserRepository {
	return &UserRepository{
		Conn: conn,
	}
}

// SignUp inserts user data into mysql
func (h *UserRepository) SignUp(d *repository.SignUp) int {

	stmt, err := h.Conn.Prepare("INSERT INTO users(name, display_name, email, password) VALUES(?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}

	pwd := []byte(d.Password)
	hashedPassword := crypto.HashAndSalt(pwd)

	res, err := stmt.Exec(d.Name, d.DisplayName, d.Email, hashedPassword)
	if err != nil {
		log.Fatal(err)
	}
	lastID, err := res.LastInsertId() // 挿入した行のIDを返却
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected() // 影響を受けた行数
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("ID = %d, affected = %d\n", lastID, rowCnt)

	return int(lastID)

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

	//return 3
}
