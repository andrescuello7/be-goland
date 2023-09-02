package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// var DSN = "host=localhost user=postgres password=251200 dbname=postgres port=5432"
var DSN = "postgresql://postgres:9c8vFku2Yxl6qypeFweC@containers-us-west-154.railway.app:5668/railway"
var DB *gorm.DB

func PostgresConnect() {
	var error error
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if error != nil {
		log.Fatal("Error connecting")
	} else {
		log.Println("Database connection successful")
	}
}
