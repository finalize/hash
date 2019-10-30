package interfaces

import (
	"database/sql"
	"log"

	"github.com/shgysd/hash/api/repository"
	"github.com/shgysd/hash/api/types"
	"github.com/shgysd/hash/api/utils/crypto"
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

// SignUp create a user
func (h *UserRepository) SignUp(b *types.SignUp) int64 {
	stmt, err := h.Conn.Prepare("INSERT INTO users(name, display_name, email, password) VALUES(?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}

	pwd := []byte(b.Password)
	hashedPassword := crypto.HashAndSalt(pwd)

	res, err := stmt.Exec(b.Name, b.DisplayName, b.Email, hashedPassword)
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

	return lastID
}

// SignIn user login
func (h *UserRepository) SignIn(b *types.SignIn) int {
	var (
		id       int
		name     string
		password string
	)
	rows, err := h.Conn.Query("SELECT id, name, password FROM users WHERE id = ?", b.ID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name, &password)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	pwd := []byte(b.Password)

	if _, err := crypto.ComparePasswords(password, pwd); err != nil {
		log.Println(err.Error())
		return 0
	}

	return id
}
