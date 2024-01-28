package db

import (
	"morseme/server/api/restricted"
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
	Id             int       `json:"id"`
	Message        string    `json:"message"`
	Sender         string    `json:"sender"`
	Ticket         string    `json:"ticket"`
	Submitted      time.Time `json:"submitted"`
	Delivered      time.Time `json:"delivered"`
	DeliveredState bool      `json:"delivered_state"`
}

const DATABASE_PATH = "app.db"

var MESSAGE_STATS_CACHE struct {
	total       int
	undelivered int
	delivered   int
}

func InitDb() {
	db, err := gorm.Open(sqlite.Open(DATABASE_PATH), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Message{})

	UpdateMessageCountCache()
}

func CreateMessage(message string, sender string, ticket string, submitted time.Time) {
	db, err := gorm.Open(sqlite.Open(DATABASE_PATH), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.Create(&Message{
		Message:   message,
		Sender:    sender,
		Ticket:    ticket,
		Submitted: submitted,
	})
}

func CreateUser(username string, password string) {
	db, err := gorm.Open(sqlite.Open(DATABASE_PATH), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	hp := restricted.HastString(password)

	db.Create(&User{
		Username: username,
		Password: hp,
	})
}

func UpdateUser(username string, password string) {
	db, err := gorm.Open(sqlite.Open(DATABASE_PATH), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	hp := restricted.HastString(password)

	var user User

	db.Where("username = ?", username).First(&user)

	user.Password = hp

	db.Save(&user)
}

func UpdateMessageCountCache() {
	db, err := gorm.Open(sqlite.Open(DATABASE_PATH), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	var messages []Message

	result := db.Find(&messages)
	total := result.RowsAffected

	result = db.Where("delivered_state = ?", 0).Find(&messages)
	undelivered := result.RowsAffected

	delivered := total - undelivered

	MESSAGE_STATS_CACHE.total = int(total)
	MESSAGE_STATS_CACHE.undelivered = int(undelivered)
	MESSAGE_STATS_CACHE.delivered = int(delivered)
}

func ReadMessageCountCache() (int, int, int) {
	return MESSAGE_STATS_CACHE.total, MESSAGE_STATS_CACHE.undelivered, MESSAGE_STATS_CACHE.delivered
}

func ReadMessage(ticket string) (Message, error) {
	db, err := gorm.Open(sqlite.Open(DATABASE_PATH), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	var message Message

	result := db.Where("ticket = ?", ticket).Last(&message)

	if result.Error != nil {
		return Message{}, result.Error
	}

	return message, nil
}

func ReadFirstUndeliveredMessage() (Message, error) {
	db, err := gorm.Open(sqlite.Open(DATABASE_PATH), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	var message Message

	result := db.Where("delivered_state = ?", 0).First(&message)

	if result.Error != nil {
		return Message{}, result.Error
	}

	return message, nil
}

func ReadLatestMessage() (Message, error) {
	db, err := gorm.Open(sqlite.Open(DATABASE_PATH), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	var message Message

	result := db.Last(&message)

	if result.Error != nil {
		return Message{}, result.Error
	}

	return message, nil
}

func ReadAllMessages() []Message {
	db, err := gorm.Open(sqlite.Open(DATABASE_PATH), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	var messages []Message

	db.Find(&messages)

	return messages
}

func ReadAllUsers() []User {
	db, err := gorm.Open(sqlite.Open(DATABASE_PATH), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	var users []User

	db.Find(&users)

	return users
}

func ReadAllUsersMap() map[string]string {
	users := ReadAllUsers()
	m := map[string]string{}

	for _, v := range users {
		m[v.Username] = v.Password
	}

	return m
}
