package service

type Users struct {
	UserID       uint `gorm:"primary_key"`
	UserName     string
	Email        string
	EmailConfirm bool
	UserStatus   string
	LoginAttempt int
	LastLogin    string
	IpAddress    string
	RegisterAt   string
	AcceptTerms  bool
}

type UserRegisterForm struct {
	UserName  string `json:"user_name"`
	Email     string `json:"email"`
	IpAddress string `json:"ip_address"`
	Password  string `json:"password"`
}

type UserAccpetTermConditionForm struct {
	UserID uint `json:"user_id"`
	Accept bool `json:"accept"`
}

type UserService interface {
	GetAllUsers() ([]Users, error)
	CreateUser(userForm *UserRegisterForm) (*Users, error)
	AcceptTermCondition(userForm *UserAccpetTermConditionForm) (bool, error)
	// GetOneUser(user *Users) (Users, error)
	// RegisterUser(*UserRegisterForm) error
}
