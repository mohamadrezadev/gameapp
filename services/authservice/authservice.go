package authservice

import (
	"GameApp/entity"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Service struct {
	signKey               string
	accessExpirationTime  time.Duration
	refreshExpirationTime time.Duration
	accessSubject         string
	refreshSubject        string
}

func New(signKey, accessSubject, refreshSubject string,
	accessExpirationTime, refreshExpirationTime time.Duration) Service{
		return Service{
			signKey: signKey,
			accessExpirationTime: accessExpirationTime,
			refreshExpirationTime: refreshExpirationTime,
			accessSubject: accessSubject,
			refreshSubject: refreshSubject,
		}
	}

func (s Service) CreateAccessToken(user entity.User)(string,error){
	return s.createToken(user.ID,s.accessSubject,s.accessExpirationTime)
}

func (s Service) CreateRefreshToken (user entity.User) (string ,error){
	return s.createToken(user.ID,s.refreshSubject,s.refreshExpirationTime)
} 

func (s Service) ParseToken(bearertoekn string)(*Claims,error){
	tokenStr:=strings.Replace(bearertoekn,"Bearer ","",1)
	token,err:=jwt.ParseWithClaims(tokenStr,&Claims{},func(t *jwt.Token) (interface{}, error) {
		return []byte(s.signKey),nil
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
	tokenstring,err:=accessToken.SignedString([]byte(s.signKey))
	if err !=nil{
		return "",err
	}
	return tokenstring,nil
	 

 }
