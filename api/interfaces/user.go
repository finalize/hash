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

// NewUserRepo returns user repository that contains db
func NewUserRepo(conn *sql.DB) repository.UserRepository {
	return &UserRepository{
		Conn: conn,
	}
}

// SignUp inserts user data into mysql and returns JWT
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
}

// SignIn returns JWT
func (h *UserRepository) SignIn(d *repository.SignIn) int {
	var (
		id       int
		password string
	)
	rows, err := h.Conn.Query("SELECT id, password FROM users WHERE id = ?", d.ID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &password)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(password)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	pwd := []byte(d.Password)

	err = crypto.ComparePasswords(password, pwd)
	if err != nil {
		log.Fatal(err)
	}

	return id
}
