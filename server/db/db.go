package db

import (
	"morseme/server/message"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Users struct {
	Id       int
	Username string
	Password string
}

type Messages struct {
	Id             int
	Message        string
	Sender         string
	Ticket         string
	Submitted      time.Time
	Delivered      time.Time
	DeliveredState bool
}

const DATABASE_PATH = "app.db"

func InitDb() {
	db, err := gorm.Open(sqlite.Open(DATABASE_PATH), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Users{})
	db.AutoMigrate(&Messages{})
}

func WriteMessage(newmsg message.Message, ticket string) {
	db, err := gorm.Open(sqlite.Open(DATABASE_PATH), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.Create(&Messages{
		Message:   newmsg.MessageText,
		Sender:    newmsg.MessageSender,
		Ticket:    ticket,
		Submitted: time.Now(),
	})
}
