package service

type LoginDefaultForm struct {
	Email    string `json:"email"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
type ResponseLogin struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

type AuthService interface {
	DefaultLogin(LoginDefaultForm *LoginDefaultForm) (*ResponseLogin, error)
}
