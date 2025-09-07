package userhandler

import (
	"GameApp/services/authservice"
	"GameApp/services/userservice"
	"GameApp/validator/uservalidator"
)

type Handler struct {
	authserv authservice.Service
	userserv userservice.Service
	userValidator uservalidator.Validator
}

func New(authserv authservice.Service,userserv userservice.Service,uservalidator uservalidator.Validator) Handler{
	return Handler{
		authserv: authserv,
		userserv: userserv,
		userValidator: uservalidator,
	}
}