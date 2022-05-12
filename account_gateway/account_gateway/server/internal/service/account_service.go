package service

import (
	"identity_provider/internal/errs"
	"identity_provider/internal/repository"

	"identity_provider/internal/utils"
	"strconv"

	"github.com/gofrs/uuid"

	"gorm.io/gorm"
)

type accountService struct {
	accountRepo repository.AccountRepository
}

func NewAccountService(accountRepo repository.AccountRepository) AccountService {
	return accountService{accountRepo: accountRepo}
}

func (a accountService) CreateAccount(createAccountForm *CreateAccountForm) (*AccountResponse, error) {

	_, err := a.accountRepo.FindAccount(repository.Account{
		UserID: createAccountForm.UserID,
	})
	if err != gorm.ErrRecordNotFound {
		return nil, errs.CustomError(
			"Account already exists Or Could not find user",
			400,
		)
	}

	if createAccountForm.MobilePhone[0] == '0' {
		createAccountForm.MobilePhone = utils.TrimFirstRune(createAccountForm.MobilePhone)
	}

	generateUUID, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	user_id := strconv.Itoa(createAccountForm.UserID)
	GenerateAccountCode := utils.Zfill(user_id, "0", 6)

	account := &repository.Account{
		UserID:         createAccountForm.UserID,
		PublicID:       generateUUID,
		ProfilePicture: createAccountForm.ProfilePicture,
		AvatarIcon:     createAccountForm.AvatarIcon,
		AvatarColor:    createAccountForm.AvatarColor,
		AvatarUsage:    createAccountForm.AvatarUsage,
		AccountCode:    "A" + GenerateAccountCode,
		NamePrefix:     createAccountForm.NamePrefix,
		MobilePrefix:   createAccountForm.MobilePrefix,
		MobilePhone:    createAccountForm.MobilePhone,
		FirstName:      createAccountForm.FirstName,
		LastName:       createAccountForm.LastName,
		DateOfBirth:    createAccountForm.DateOfBirth,
		GenderID:       createAccountForm.GenderID,
		CountryID:      createAccountForm.CountryID,
	}

	acc, err := a.accountRepo.CreateAccount(*account)
	if err != nil {
		return nil, err

	}
	return &AccountResponse{
		AccountID:      acc.AccountID,
		UserID:         acc.UserID,
		PublicID:       acc.PublicID,
		ProfilePicture: acc.ProfilePicture,
		AvatarIcon:     acc.AvatarIcon,
		AvatarColor:    acc.AvatarColor,
		AvatarUsage:    acc.AvatarUsage,
		LocalLanguage:  acc.LocalLanguage,
		NightMode:      acc.NightMode,
		AccountCode:    acc.AccountCode,
		AccountGroupID: acc.AccountGroupID,
		NamePrefix:     acc.NamePrefix,
		MobilePrefix:   acc.MobilePrefix,
		MobilePhone:    acc.MobilePhone,
		FirstName:      acc.FirstName,
		LastName:       acc.LastName,
		DateOfBirth:    acc.DateOfBirth,
		Timezone:       acc.Timezone,
		GenderID:       acc.GenderID,
		CountryID:      acc.CountryID,
	}, nil

}

func (a accountService) FindAccount(RequstFindAccount *RequstFindAccount) (*AccountResponse, error) {

	acc, err := a.accountRepo.FindAccount(repository.Account{
		UserID: RequstFindAccount.UserID,
	})

	if err == gorm.ErrRecordNotFound {
		return nil, errs.CustomError(
			"Account already exists Or Could not find user",
			400,
		)
	}
	return &AccountResponse{
		AccountID:      acc.AccountID,
		UserID:         acc.UserID,
		PublicID:       acc.PublicID,
		ProfilePicture: acc.ProfilePicture,
		AvatarIcon:     acc.AvatarIcon,
		AvatarColor:    acc.AvatarColor,
		AvatarUsage:    acc.AvatarUsage,
		LocalLanguage:  acc.LocalLanguage,
		NightMode:      acc.NightMode,
		AccountCode:    acc.AccountCode,
		AccountGroupID: acc.AccountGroupID,
		NamePrefix:     acc.NamePrefix,
		MobilePrefix:   acc.MobilePrefix,
		MobilePhone:    acc.MobilePhone,
		FirstName:      acc.FirstName,
		LastName:       acc.LastName,
		DateOfBirth:    acc.DateOfBirth,
		Timezone:       acc.Timezone,
		GenderID:       acc.GenderID,
		CountryID:      acc.CountryID,
	}, nil

}

func (a accountService) UpdateAccount(UpdateAccountForm *UpdateAccountForm) (*AccountResponse, error) {

	_, err := a.accountRepo.FindAccount(repository.Account{
		UserID: UpdateAccountForm.UserID,
	})
	if err == gorm.ErrRecordNotFound {
		return nil, errs.CustomError(
			"Account already exists Or Could not find user",
			400,
		)
	}

	if UpdateAccountForm.MobilePhone[0] == '0' {
		UpdateAccountForm.MobilePhone = utils.TrimFirstRune(UpdateAccountForm.MobilePhone)
	}

	accountUserID := &repository.Account{
		UserID:         UpdateAccountForm.UserID,
		ProfilePicture: UpdateAccountForm.ProfilePicture,
		AvatarIcon:     UpdateAccountForm.AvatarIcon,
		AvatarColor:    UpdateAccountForm.AvatarColor,
		AvatarUsage:    UpdateAccountForm.AvatarUsage,
		LocalLanguage:  UpdateAccountForm.LocalLanguage,
		NightMode:      UpdateAccountForm.NightMode,
		NamePrefix:     UpdateAccountForm.NamePrefix,
		MobilePrefix:   UpdateAccountForm.MobilePrefix,
		MobilePhone:    UpdateAccountForm.MobilePhone,
		FirstName:      UpdateAccountForm.FirstName,
		LastName:       UpdateAccountForm.LastName,
		DateOfBirth:    UpdateAccountForm.DateOfBirth,
		Timezone:       UpdateAccountForm.Timezone,
		GenderID:       UpdateAccountForm.GenderID,
		CountryID:      UpdateAccountForm.CountryID,
	}

	a.accountRepo.UpdateOneAccount(accountUserID, accountUserID.UserID)

	acc, err := a.accountRepo.FindAccount(repository.Account{
		UserID: UpdateAccountForm.UserID,
	})

	if err == gorm.ErrRecordNotFound {
		return nil, errs.CustomError(
			"Account already exists Or Could not find user",
			400,
		)
	}
	return &AccountResponse{
		AccountID:      acc.AccountID,
		UserID:         acc.UserID,
		PublicID:       acc.PublicID,
		ProfilePicture: acc.ProfilePicture,
		AvatarIcon:     acc.AvatarIcon,
		AvatarColor:    acc.AvatarColor,
		AvatarUsage:    acc.AvatarUsage,
		LocalLanguage:  acc.LocalLanguage,
		NightMode:      acc.NightMode,
		AccountCode:    acc.AccountCode,
		AccountGroupID: acc.AccountGroupID,
		NamePrefix:     acc.NamePrefix,
		MobilePrefix:   acc.MobilePrefix,
		MobilePhone:    acc.MobilePhone,
		FirstName:      acc.FirstName,
		LastName:       acc.LastName,
		DateOfBirth:    acc.DateOfBirth,
		Timezone:       acc.Timezone,
		GenderID:       acc.GenderID,
		CountryID:      acc.CountryID,
	}, nil

}
