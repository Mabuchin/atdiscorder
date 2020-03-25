package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var db *gorm.DB

func InitDB() *gorm.DB {
	var err error
	db, err = gorm.Open("sqlite3", "problem_list.db")
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(&Problem{})
	return db
}
