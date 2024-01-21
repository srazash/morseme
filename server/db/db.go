package db

import (
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

func InsertMessage(message string, sender string, ticket string, submitted time.Time) {
	db, err := gorm.Open(sqlite.Open(DATABASE_PATH), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.Create(&Messages{
		Message:   message,
		Sender:    sender,
		Ticket:    ticket,
		Submitted: submitted,
	})
}

func CountMessages() (int64, int64, int64) {
	db, err := gorm.Open(sqlite.Open(DATABASE_PATH), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	var messages []Messages

	result := db.Find(&messages)

	total := result.RowsAffected

	result = db.Where("delivered_state = ?", 0).Find(&messages)

	undelivered := result.RowsAffected

	delivered := total - undelivered

	return total, undelivered, delivered
}
