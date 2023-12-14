package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	Ticket  string
	Message string
	Status  bool
}

func Connect() {
	db, err := gorm.Open(sqlite.Open("msg.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Message{})
}
