package model

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func AutoMigrate() {
	_, err := os.Stat("test.db")
	if err == nil {
		return
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Todo{})

	db.Create(&User{ID: "demo", Name: "demo"})
}
