package service

import (
	"identity_provider/internal/errs"
	"identity_provider/internal/repository"
	"identity_provider/internal/utils"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return userService{userRepo: userRepo}
}

func (s userService) GetAllUsers() ([]Users, error) {
	users, err := s.userRepo.GetAll()
	if err != nil {
		return nil, err
	}
	userResponses := []Users{}
	for _, user := range users {
		userResponse := Users{
			UserID:       (user.UserID),
			UserName:     user.UserName,
			Email:        user.Email,
			EmailConfirm: user.EmailConfirm,
			UserStatus:   user.UserStatus,
			LoginAttempt: user.LoginAttempt,
			LastLogin:    user.LastLogin,
			IpAddress:    user.IpAddress,
			RegisterAt:   user.RegisterAt,
			AcceptTerms:  user.AcceptTerms,
		}
		userResponses = append(userResponses, userResponse)
	}
	return userResponses, nil
}

func (s userService) CreateUser(userForm *UserRegisterForm) (*Users, error) {

	if len(userForm.UserName) < 5 || len(userForm.UserName) > 20 {
		return nil, errs.CustomError(
			"Username must be more than 5 and less than 20 characters",
			400,
		)
	}

	userInputPassword := userForm.Password
	sevenOrMore, number, upper, special := utils.VerifyPassword(userInputPassword)
	if !sevenOrMore || !number || !upper || !special {
		return nil, errs.CustomError(
			"Password must be at least 7 characters long, contain at least one number, one upper case letter, and one special character",
			400,
		)
	}

	if !utils.VerifyEmail(userForm.Email) {
		return nil, errs.CustomError(
			"Email is not valid",
			400,
		)
	}

	_, err := s.userRepo.FindUser(
		repository.User{
			Email:    userForm.Email,
			UserName: userForm.UserName,
		})

	if err != gorm.ErrRecordNotFound {
		return nil, errs.CustomError(
			"Email or Username is already taken",
			400,
		)
	}

	password, err := bcrypt.GenerateFromPassword([]byte(userInputPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, errs.CustomError(
			"Error while hashing password",
			500,
		)
	}
	userRegister := repository.User{
		UserName:     userForm.UserName,
		Email:        userForm.Email,
		HashPassword: string(password),
		IpAddress:    userForm.IpAddress,
		RegisterAt:   strconv.Itoa(int(time.Now().Unix())),
		UserStatus:   "active",
	}
	user, err := s.userRepo.Create(userRegister)
	if err != nil {
		return nil, errs.CustomError(
			"Error while creating user",
			500,
		)
	}

	if user != nil {
		userResponse := Users{
			UserID:       (user.UserID),
			UserName:     user.UserName,
			Email:        user.Email,
			EmailConfirm: user.EmailConfirm,
			UserStatus:   user.UserStatus,
			LoginAttempt: user.LoginAttempt,
			LastLogin:    user.LastLogin,
			IpAddress:    user.IpAddress,
			RegisterAt:   user.RegisterAt,
		}

		return &userResponse, nil

	}
	return nil, errs.CustomError(
		"Error while creating user",
		500,
	)
}

func (s userService) AcceptTermCondition(userForm *UserAccpetTermConditionForm) (bool, error) {

	if userForm.UserID == 0 {
		return false, errs.CustomError(
			"User ID is required",
			400,
		)
	}

	if !userForm.Accept {
		return false, errs.CustomError(
			"Accept term and condition is required",
			400,
		)
	}

	user, err := s.userRepo.GetById((int(userForm.UserID)))
	if err == gorm.ErrRecordNotFound {
		return false, errs.CustomError(
			"Not Found This UserName",
			400,
		)
	}

	if user.AcceptTerms {
		return false, errs.CustomError(
			"User already accepted term and condition",
			400,
		)

	}

	_, err = s.userRepo.UpdateOne(repository.User{
		UserID:      user.UserID,
		AcceptTerms: userForm.Accept,
	})

	if err != nil {
		return false, errs.CustomError(
			"Error while updating user",
			500,
		)

	}
	return true, nil
}
