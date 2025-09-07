package uservalidator

import "GameApp/entity"


const (
	phoneNumberRegex = "^09[0-9]{9}$"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	GetUserByPhoneNumber(phoneNumber string) (entity.User, error)
}

type Validator struct {
	repository Repository
}

func New(repository Repository) Validator {
	return Validator{repository: repository}
}
