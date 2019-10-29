package types

type SignUp struct {
	Name        string `json:"name" validate:"required"`        //必須パラメータ
	DisplayName string `json:"displayName" validate:"required"` //必須パラメータ
	Email       string `json:"email" validate:"required,email"` //必須パラメータ、かつ、emailフォーマット
	Password    string `json:"password" validate:"required"`
}

type SignIn struct {
	ID       string `json:"id" validate:"required"`
	Password string `json:"password" validate:"required"`
}
