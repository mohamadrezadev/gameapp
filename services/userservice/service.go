package userservice

import (
	"GameApp/entity"
	"crypto/md5"
	"encoding/hex"
	"time"
	jwt "github.com/golang-jwt/jwt/v4"
)

type Repository interface {
	IsPhoneNumberUnique(PhoneNumber string) (bool, error)
	RegisterUser(u entity.User) (entity.User, error)
	GetUserByPhoneNumber(PhoneNumber string) (entity.User, error)
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


func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
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
