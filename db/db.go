package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func Init() (*sql.DB, error) {
	cfg := mysql.Config{
		User:   viper.GetString("database.user"),
		Passwd: viper.GetString("database.pass"),
		Net:    "tcp",
		Addr:   fmt.Sprintf("%s:%s", viper.GetString("database.host"), viper.GetString("database.port")),
		DBName: viper.GetString("database.name"),
	}
	var err error
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	return db, err
}
