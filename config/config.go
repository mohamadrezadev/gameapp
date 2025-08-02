package config

import (
	"GameApp/repository/mysql"
	"GameApp/services/authservice"
)

type HTTPSrver struct {
	Port int
}

type Config struct {
	HTTPServer HTTPSrver
	Auth authservice.Config
	Mysql      mysql.Config
}