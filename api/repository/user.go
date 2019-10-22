package repository

import (
	"github.com/shgysd/hash/api/model"
)

// UserRepository Define user method
type UserRepository interface {
	SignUp(u *model.User)
}
