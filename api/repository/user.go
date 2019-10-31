package repository

// SignUp contains default user info
type SignUp struct {
	Name        string `json:"name" validate:"required"`
	DisplayName string `json:"displayName" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required"`
}

// UserRepository defines user method
type UserRepository interface {
	SignUp(d *SignUp) int
}
