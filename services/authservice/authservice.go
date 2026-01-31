package authservice

import (
	"GameApp/entity"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Config struct {
	SignKey               string `koanf:"signkey"`
	AccessExpirationTime  time.Duration `koanf:"access_expiration_time"`
	RefreshExpirationTime time.Duration `koanf:"refresh_expiration_time"`
	AccessSubject         string `koanf:"access_subject"`
	RefreshSubject        string `koanf:"refresh_subject"`
}
type Service struct{
	config  Config
}

func New(cfg Config) Service{
		return Service{
			config:cfg,
		}
	}

func (s Service) CreateAccessToken(user entity.User)(string,error){
	return s.createToken(user.ID,s.config.AccessSubject,s.config.AccessExpirationTime)
}

func (s Service) CreateRefreshToken (user entity.User) (string ,error){
	return s.createToken(user.ID,s.config.RefreshSubject,s.config.RefreshExpirationTime)
} 

func (s Service) ParseToken(bearertoekn string)(*Claims,error){
	tokenStr:=strings.Replace(bearertoekn,"Bearer ","",1)
	token,err:=jwt.ParseWithClaims(tokenStr,&Claims{},func(t *jwt.Token) (interface{}, error) {
		return []byte(s.config.SignKey),nil
	})
	
	if err!=nil{
		return nil,err
	}

	if claims,ok := token.Claims.(*Claims);ok && token.Valid{
		return claims,nil
	}else{
		return nil,err
	}
}
func (s Service) createToken(userId uint,subject string,expireDuration  time.Duration) (string ,error){
 	claim:=Claims{
		RegisteredClaims:jwt.RegisteredClaims{
			Subject: subject,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireDuration)),
		},
		UserID:userId,
	}

	accessToken:=jwt.NewWithClaims(jwt.SigningMethodES256,claim)
	tokenstring,err:=accessToken.SignedString([]byte(s.config.SignKey))
	if err !=nil{
		return "",err
	}
	return tokenstring,nil
	 

 }
