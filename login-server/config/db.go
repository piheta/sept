package config

import (
	"fmt"
	"os"

	"github.com/piheta/sept/login-server/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// FOR DEV
// docker run --name sept-login-server -p 5432:5432 -e POSTGRES_PASSWORD=password -d postgres
// export DATABASE_SECRET="host=localhost user=postgres password=password dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Oslo"

func Connect() {
	dsn := os.Getenv("DATABASE_SECRET")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	db.AutoMigrate(&models.User{})

	DB = db
}
