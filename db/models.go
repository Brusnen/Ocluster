package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Runner struct {
	gorm.Model
	Name    string
	Address string
	IsAlive bool
}

func InitDb() {
	db, err := gorm.Open(sqlite.Open("db.db"), &gorm.Config{})
	if err != nil {
		panic("Error init database")
	}
	db.AutoMigrate(&Runner{})
}
