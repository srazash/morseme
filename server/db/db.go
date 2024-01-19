package db

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Message struct {
	Id             int
	Message        string
	Sender         string
	Ticket         string
	Submitted      time.Time
	Delivered      time.Time
	DeliveredState bool
}

func Connect() {
	db, err := gorm.Open(sqlite.Open("msg.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Message{})
}
