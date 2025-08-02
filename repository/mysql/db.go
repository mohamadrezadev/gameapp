package mysql

import (
	"database/sql"
	"fmt"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Username string
	Password string
	Port     int
	Host     string
	DBName   string
}
type MySqlDb struct {
	Config Config
	db     *sql.DB
}

func New(config Config) *MySqlDb{
		db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true",
		config.Username, config.Password, config.Host, config.Port, config.DBName))
	
	if err!=nil{
		panic(fmt.Errorf("cant open mysql db:%v ",err))
	}
	db.SetConnMaxLifetime(time.Minute *3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &MySqlDb{Config: config,db: db}
}