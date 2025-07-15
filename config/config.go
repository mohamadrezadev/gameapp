package config

import "GameApp/repository/mysql"

type HTTPSrver struct {
	Port int
}

type Config struct {
	HTTPServer HTTPSrver
	Mysql      mysql.Config
}