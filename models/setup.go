package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDatabase() {

	// viper.AutomaticEnv() // To get the value from the config file using key// viper package read .env
	// viper_user := viper.Get("POSTGRES_USER")
	// viper_password := viper.Get("POSTGRES_PASSWORD")
	// viper_db := viper.Get("POSTGRES_DB")
	// viper_host := viper.Get("POSTGRES_HOST")
	// viper_port := viper.Get("POSTGRES_PORT") // https://gobyexample.com/string-formatting

	viper_user := "postgres"
	viper_password := "postgres"
	viper_db := "postgres"
	viper_host := "localhost"
	viper_port := 5432

	prosgret_conname := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", viper_host, viper_port, viper_user, viper_db, viper_password)
	fmt.Println("conname is\t\t", prosgret_conname)
	db, err := gorm.Open("postgres", prosgret_conname)

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Book{}) // Initialise value
	// m := Book{Author: "author1", Title: "title1"}

	// db.Create(&m)

	DB = db
}
