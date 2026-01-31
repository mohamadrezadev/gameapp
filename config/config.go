package config

import (
	"GameApp/repository/mysql"
	"GameApp/services/authservice"
)

type HTTPSrver struct {
	Port int `koanf:"port"`
}

type Config struct {
	HTTPServer HTTPSrver `koanf:"http_server"`
	Auth authservice.Config `koanf:"auth"`
	Mysql      mysql.Config `koanf:"mysql"`
}