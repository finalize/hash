package structs

type SignIn struct {
	ID       int    `json:"id" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type GetUser struct {
	ID int `json:"id"`
}
