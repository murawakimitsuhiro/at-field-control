package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
	"log"
	"time"
)

var DB *gorm.DB

type Model struct {
	ID        uint64     `json:"id" gorm:"primary_key"`
	createdat *time.Time `json:"-"`
	updatedat *time.Time `json:"-"`
	deletedat *time.Time `json:"-" sql:"index"`
}

func InitDB() {
	db, err := gorm.Open("mysql", dbInfoString())
	if err != nil {
		log.Panicf("Failed to connect to database: %v\n", err)
	}

	db.DB()
	db.AutoMigrate(&User{})

	DB = db

}

func dbInfoString() string {
	s, b := revel.Config.String("db.info")
	if !b {
		log.Panicf("database info not found")
	}

	return s
}
