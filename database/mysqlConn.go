package database

import (
	"log"

	"github.com/aabdullahgungor/golang-jwt-restapi/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	DB_DRIVER   = "mysql"
	DB_HOST = "localhost:3306"
	DB_NAME     = "jwt_demo"
	DB_USER     = "root"
	DB_PASSWORD = "mysql123"
)
var db *gorm.DB
var errDb error

func Connect()  {

	db, errDb = gorm.Open(mysql.Open(DB_USER + ":" + DB_PASSWORD + "@tcp(" + DB_HOST + ")/" + DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if errDb != nil {
		log.Fatal(errDb)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database!")
}

func Migrate() {
	db.AutoMigrate(&model.User{})
	log.Println("Database Migration Completed!")
}

