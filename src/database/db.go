package database

import (
	"ambassador/src/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func Connect () {
	var err error
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		dbName,
		os.Getenv("DB_LOC"))

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//_, err := gorm.Open(mysql.Open("root:root@tcp(db:3306)/ambassador"), &gorm.Config{})

	if err != nil {
		panic("Could not connect with the database!")
	}

}

func AutoMigrate(){
	// 以下のように書くと、gormがUser structを取得し、DBにtableを作成してくれる
	DB.AutoMigrate(models.User{})
}