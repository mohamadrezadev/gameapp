package userservice

import (
	"GameApp/param"
	"GameApp/pkg/richerror"
)

func (s Service) Profile(req param.ProfileRequest) (param.ProfileResponse, error) {
	const op = "userservice.Profile"
	user, err := s.repo.GetUserById(req.UserID)
	if err != nil {
		return param.ProfileResponse{}, richerror.New(op).WithErr(err).
			WithMeta(map[string]interface{}{"req": req})
	}
	return param.ProfileResponse{Name: user.Name}, nil
}
