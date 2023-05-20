package main

import (
	"Projekat/database"
	"database/sql"
	"fmt"
	"log"

	"github.com/go-ini/ini"
	_ "github.com/go-sql-driver/mysql"
)

type MySQLConfig struct {
	Username string
	Password string
	Database string
}

func main() {
	cfg, err := ini.Load("database_information.ini")
	if err != nil {
		log.Println(err)
		return
	}

	mysqlSection := cfg.Section("MySQL")
	if mysqlSection == nil {
		log.Println("MySQL section not found in the configuration file")
		return
	}

	mysqlConfig := MySQLConfig{}
	err = mysqlSection.MapTo(&mysqlConfig)
	if err != nil {
		log.Println(err)
		return
	}

	database.Db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", mysqlConfig.Username, mysqlConfig.Password, mysqlConfig.Database))
	if err != nil {
		log.Fatal(err)
	}
	defer database.Db.Close()

	err = database.Db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	httpServer(database.Db)
}
