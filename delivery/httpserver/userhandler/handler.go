package userhandler

import (
	"GameApp/services/authservice"
	"GameApp/services/userservice"
	"GameApp/validator/uservalidator"
)

type Handler struct {
	authconfig authservice.Config
	authserv authservice.Service
	userserv userservice.Service
	userValidator uservalidator.Validator
}

func New(authconfig authservice.Config, authserv authservice.Service, userserv userservice.Service, uservalidator uservalidator.Validator) Handler{
	return Handler{
		authconfig: authconfig,
		authserv: authserv,
		userserv: userserv,
		userValidator: uservalidator,
	}
}