package repository

import "github.com/shgysd/hash/api/types"

// UserRepository Define user method
type UserRepository interface {
	SignUp(b *types.SignUp) int64
	SignIn(b *types.SignIn) int
}
