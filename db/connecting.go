package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "root"
	password = "root"
	dbname   = "medicines"
)

const DSN = "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=disable"

var DB *gorm.DB

func Connection() {

	var error error

	DB, error = gorm.Open(
		postgres.Open(DSN),
		&gorm.Config{},
	)

	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("DB Connected")
	}
}
