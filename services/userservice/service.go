package userservice

import (
	"GameApp/dto"
	"GameApp/entity"
	"GameApp/pkg/richerror"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

type Repository interface {
	IsPhoneNumberUnique(PhoneNumber string) (bool, error)
	RegisterUser(u entity.User) (entity.User, error)
	GetUserByPhoneNumber(PhoneNumber string) (entity.User, bool, error)
	GetUserById(userId uint) (entity.User, error)
}

type AuthGenerator interface {
	CreateAccessToken(user entity.User) (string, error)
	CreateRefreshToken(user entity.User) (string, error)
}

type Service struct {
	auth AuthGenerator
	repo Repository
}

func New(repository Repository, auth AuthGenerator) Service {
	return Service{repo: repository, auth: auth}
}

func (s Service) Register(req dto.RegisterRequest) (dto.RegisterResponse, error) {

	// TODO - replace md5 with bcrypt
	user := entity.User{
		ID:          0,
		PhoneNumber: req.PhoneNumber,
		Name:        req.Name,
		Password:    getMD5Hash(req.Password),
	}
	createduser, err := s.repo.RegisterUser(user)
	if err != nil {
		return dto.RegisterResponse{}, fmt.Errorf("unexpected error: %w", err)

	}
	return dto.RegisterResponse{User: dto.UserInfo{
		ID:          createduser.ID,
		PhoneNumber: createduser.Name,
		Name:        createduser.PhoneNumber,
	}}, nil

}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type LoginResponse struct {
	AccsessToken string       `json:"access_token"`
	RefreshToken string       `json:"refresh_token"`
	User         dto.UserInfo `json:"user"`
	// Tokens Tokens       `json:"tokens"`
}

func (s Service) Login(req LoginRequest) (LoginResponse, error) {

	const op = "userservice.Login"

	user, exist, err := s.repo.GetUserByPhoneNumber(req.PhoneNumber)
	if err != nil {
		return LoginResponse{}, richerror.New(op).WithErr(err).
			WithMeta(map[string]interface{}{"phone_number": req.PhoneNumber})
	}
	if !exist {
		return LoginResponse{}, fmt.Errorf("username or passwprd is not currect ")
	}

	if user.Password != getMD5Hash(req.Password) {
		return LoginResponse{}, fmt.Errorf("username or passwprd is not currect ")

	}
	accessToken, err := s.auth.CreateAccessToken(user)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	refreshToken, err := s.auth.CreateRefreshToken(user)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	return LoginResponse{AccsessToken: accessToken, RefreshToken: refreshToken}, nil
}

type ProfileRequest struct {
	UserId uint
}

type ProfileResponse struct {
	Name string `json:"name"`
}

func (s Service) Profile(req ProfileRequest) (ProfileResponse, error) {
	const op = "userservice.Profile"
	user, err := s.repo.GetUserById(req.UserId)
	if err != nil {
		return ProfileResponse{}, richerror.New(op).WithErr(err).
			WithMeta(map[string]interface{}{"req": req})
	}
	return ProfileResponse{Name: user.Name}, nil
}

type Claims struct {
	RegisteredClaims jwt.RegisteredClaims
	UserId           uint
}

func (c Claims) Valid() error {
	return nil
}

func createtoken(userId uint, signKey string) (string, error) {
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
		},
		UserId: userId,
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := accessToken.SignedString([]byte(signKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil

}
