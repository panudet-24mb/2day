package service

import (
	"fmt"
	"identity_provider/internal/errs"
	"identity_provider/internal/repository"
	"identity_provider/internal/utils"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return authService{userRepo: userRepo}
}

func (a authService) DefaultLogin(loginform *LoginDefaultForm) (*ResponseLogin, error) {

	var (
		EmailExist    bool
		UserNameExist bool
	)

	if len(loginform.UserName) > 0 && len(loginform.Email) == 0 {
		if len(loginform.UserName) < 5 && len(loginform.UserName) >= 20 {
			return nil, errs.CustomError(
				"UserName must be between 5 and 20 characters",
				400,
			)

		}
		UserNameExist = true
	}
	if len(loginform.Email) > 0 && len(loginform.UserName) == 0 {
		email := utils.VerifyEmail(loginform.Email)
		if !email {
			return nil, errs.CustomError(
				"Email is not valid",
				400,
			)
		}
		EmailExist = true

	}

	sevenOrMore, number, upper, special := utils.VerifyPassword(loginform.Password)
	if !sevenOrMore || !number || !upper || !special {
		return nil, errs.CustomError(
			"Password must be at least 7 characters long, contain at least one number, one upper case letter, and one special character",
			400,
		)
	}
	//create empty struct

	var userInfomation []repository.User
	switch {
	case UserNameExist:
		user, err := a.userRepo.FindUser(repository.User{
			UserName: loginform.UserName,
		})
		if err == gorm.ErrRecordNotFound {
			return nil, errs.CustomError(
				"Not Found This UserName",
				400,
			)
		}
		userInfomation = append(userInfomation, *user)

	case EmailExist:
		user, err := a.userRepo.FindUser(repository.User{
			Email: loginform.Email,
		})
		if err == gorm.ErrRecordNotFound {
			return nil, errs.CustomError(
				"Not Found This Email",
				400,
			)
		}
		userInfomation = append(userInfomation, *user)

	default:
		return nil, errs.CustomError(
			"Only UserName Or Email must be filled Once",
			400,
		)
	}

	if userInfomation[0].LoginAttempt > viper.GetInt("login.maxattemptcount") {

		return nil, errs.CustomError(
			"You have exceeded the maximum number of login attempts. Please try again later.",
			400,
		)

	}

	err := bcrypt.CompareHashAndPassword([]byte(userInfomation[0].HashPassword), []byte(loginform.Password))
	if err != nil {

		user, err := a.userRepo.UpdateOne(repository.User{
			UserID:       userInfomation[0].UserID,
			LoginAttempt: userInfomation[0].LoginAttempt + 1,
		})
		if err != nil {
			return nil, errs.CustomError(
				"can't update user",
				500,
			)
		}
		return nil, errs.CustomError(
			"Password is not correct , Login Attemp Remaining "+fmt.Sprintf("%d", viper.GetInt("login.maxattemptcount")-user.LoginAttempt),
			400,
		)

	}

	cliams := jwt.StandardClaims{
		Issuer:    fmt.Sprint(userInfomation[0].UserID),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams)
	token, err := jwtToken.SignedString([]byte(viper.GetString("app.secret")))
	if err != nil {
		return nil, err
	}
	if len(token) > 0 {
		a.userRepo.UpdateOne(
			repository.User{
				UserID:       userInfomation[0].UserID,
				LoginAttempt: 1,
				LastLogin:    strconv.Itoa(int(time.Now().Unix())),
			},
		)

	}

	return &ResponseLogin{
		Message: token,
		Status:  true,
	}, nil

}
