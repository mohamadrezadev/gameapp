package userservice

import (
	"GameApp/param"
	"GameApp/pkg/richerror"
	"fmt"
)

func (s Service) Login(req param.LoginRequest) (param.LoginResponse, error) {

	const op = "userservice.Login"

	user, err := s.repo.GetUserByPhoneNumber(req.PhoneNumber)
	if err != nil {
		return param.LoginResponse{}, richerror.New(op).WithErr(err).
			WithMeta(map[string]interface{}{"phone_number": req.PhoneNumber})
	}

	if user.Password != getMD5Hash(req.Password) {
		return param.LoginResponse{}, fmt.Errorf("username or passwprd is not currect ")

	}
	accessToken, err := s.auth.CreateAccessToken(user)
	if err != nil {
		return param.LoginResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	refreshToken, err := s.auth.CreateRefreshToken(user)
	if err != nil {
		return param.LoginResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	return param.LoginResponse{Tokens: param.Tokens{AccessToken: accessToken, RefreshToken: refreshToken}}, nil
}
