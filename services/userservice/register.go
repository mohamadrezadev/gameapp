package userservice

import (
	"GameApp/entity"
	"GameApp/param"
	"fmt"
)

func (s Service) Register(req param.RegisterRequest) (param.RegisterResponse, error) {

	// TODO - replace md5 with bcrypt
	user := entity.User{
		ID:          0,
		PhoneNumber: req.PhoneNumber,
		Name:        req.Name,
		Password:    getMD5Hash(req.Password),
	}
	createduser, err := s.repo.RegisterUser(user)
	if err != nil {
		return param.RegisterResponse{}, fmt.Errorf("unexpected error: %w", err)

	}
	return param.RegisterResponse{User: param.UserInfo{
		ID:          createduser.ID,
		PhoneNumber: createduser.Name,
		Name:        createduser.PhoneNumber,
	}}, nil

}