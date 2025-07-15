package main

import (
	"GameApp/repository/mysql"
	"GameApp/services/authservice"
	"GameApp/services/userservice"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	JwtSignKey                 = "jwt_secret"
	AccessTokenSubject         = "ac"
	RefreshTokenSubject        = "rt"
	AccessTokenExpireDuration  = time.Hour * 24
	RefreshTokenExpireDuration = time.Hour * 24 * 7
)

var configdb = mysql.Config{
	Username: "gameapp",
	Password: "gameappRoo7t0lk2o20",
	Port:     3308,
	Host:     "http://localhost/",
	DBName:   "gameapp_db",
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health-check", healthCheckHandler)
	mux.HandleFunc("/users/Register", userRegisterHandler)
	mux.HandleFunc("/users/Login", userLoginHandler)
	mux.HandleFunc("/users/Profile", userProfileHandler)

	log.Println("server is listening  on port 7000... ")
	server := http.Server{Addr: ":7000", Handler: mux}
	log.Fatal(server.ListenAndServe())
}

func userRegisterHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		fmt.Fprintf(writer, `{"error": "invalid method"}`)
	}
	data, err := io.ReadAll(req.Body)
	if err != nil {
		writer.Write([]byte(
			fmt.Sprintf(`{"error": "%s"}`, err.Error()),
		))
	}

	var uReq userservice.RegisterRequest
	jerr := json.Unmarshal(data, &uReq)
	fmt.Println(jerr)
	if jerr != nil {
		writer.Write([]byte(
			fmt.Sprintf(`{"error": "%s"}`, err.Error()),
		))
		return
	}

	authSvc := authservice.New(JwtSignKey, AccessTokenSubject, RefreshTokenSubject,
		AccessTokenExpireDuration, RefreshTokenExpireDuration)

	mysqlRepo := mysql.New(configdb)
	userserv := userservice.New(mysqlRepo, authSvc)

	_, Rerr := userserv.Register(uReq)
	if Rerr != nil {
		writer.Write([]byte(
			fmt.Sprintf(`{"error": "%s"}`, err.Error()),
		))

		return
	}
	writer.Write([]byte(`{"message": "user created"}`))

}

func healthCheckHandler(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, `{"message": "everything is good!"}`)
}

func userLoginHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		fmt.Fprintf(writer, `{"error": "invalid method"}`)
	}
	data, err := io.ReadAll(req.Body)
	if err != nil {
		writer.Write([]byte(
			fmt.Sprintf(`{"error": "%s"}`, err.Error()),
		))
	}

	var IReq userservice.LoginRequest
	err = json.Unmarshal(data, &IReq)
	if err != nil {
		writer.Write([]byte(
			fmt.Sprintf(`{"error": "%s"}`, err.Error()),
		))

		return
	}
	authSvc := authservice.New(JwtSignKey, AccessTokenSubject, RefreshTokenSubject,
		AccessTokenExpireDuration, RefreshTokenExpireDuration)

	mysqlRepo := mysql.New(configdb)
	userSvc := userservice.New(mysqlRepo, authSvc)

	resp, err := userSvc.Login(IReq)
	if err != nil {
		writer.Write([]byte(
			fmt.Sprintf(`{"error": "%s"}`, err.Error()),
		))

		return
	}

	data, err = json.Marshal(resp)
	if err != nil {
		writer.Write([]byte(
			fmt.Sprintf(`{"error": "%s"}`, err.Error()),
		))

		return
	}
	writer.Write(data)
}

func userProfileHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		fmt.Fprintf(writer, `{"error": "invalid method"}`)
	}
	authSvc := authservice.New(JwtSignKey, AccessTokenSubject, RefreshTokenSubject,
		AccessTokenExpireDuration, RefreshTokenExpireDuration)

	authToken := req.Header.Get("Authorization")
	claims, err := authSvc.ParseToken(authToken)
	if err != nil {
		fmt.Fprintf(writer, `{"error": "token is not valid"}`)
	}

	mysqlRepo := mysql.New(configdb)
	userSvc := userservice.New(mysqlRepo, authSvc)

	resp, err := userSvc.Profile(userservice.ProfileRequest{UserId: claims.UserID})
	if err != nil {
		writer.Write([]byte(
			fmt.Sprintf(`{"error": "%s"}`, err.Error()),
		))

		return
	}

	data, err := json.Marshal(resp)
	if err != nil {
		writer.Write([]byte(
			fmt.Sprintf(`{"error": "%s"}`, err.Error()),
		))

		return
	}

	writer.Write(data)
}
