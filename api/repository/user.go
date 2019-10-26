package repository

import "github.com/shgysd/hash/api/types"

// UserRepository Define user method
type UserRepository interface {
	SignUp(j *types.SignUp) int64
}
