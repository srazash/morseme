package db

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	Id       int
	Username string
	Password string
}

type Message struct {
	Id             int
	Message        string
	Sender         string
	Ticket         string
	Submitted      time.Time
	Delivered      time.Time
	DeliveredState bool
}

func CheckDb() {
	db, err := gorm.Open(sqlite.Open("msg.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Message{})
}
