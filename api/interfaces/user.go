package interfaces

import (
	"database/sql"
	"log"

	"github.com/shgysd/hash/api/repository"
	"github.com/shgysd/hash/api/types"
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
func (h *UserRepository) SignUp(j *types.SignUp) int64 {

	stmt, err := h.Conn.Prepare("INSERT INTO users(name, display_name, email, password) VALUES(?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(j.Name, j.DisplayName, j.Email, j.Password) // => "INSERT INTO users(name) VALUES('Dolly')"
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

	// var (
	// 	id   int
	// 	name string
	// )
	// rows, err := h.Conn.Query("SELECT id, name FROM users")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer rows.Close()
	// for rows.Next() {
	// 	err := rows.Scan(&id, &name)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Println(id, name)
	// }
	// err = rows.Err()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	log.Printf("ID = %d, affected = %d\n", lastID, rowCnt)

	return lastID

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
