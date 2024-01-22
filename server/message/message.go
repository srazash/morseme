package message

import (
	"errors"
	"fmt"
	"morseme/server/db"
	"morseme/server/ticket"
	"regexp"
	"time"

	"github.com/labstack/gommon/log"
)

type Message struct {
	MessageId      int       `json:"message_id"`
	MessageText    string    `json:"message_text"`
	MessageSender  string    `json:"message_sender"`
	MessageTicket  string    `json:"message_ticket"`
	Submitted      time.Time `json:"submitted"`
	Delivered      time.Time `json:"delivered"`
	DeliveredState bool      `json:"delivered_state"`
}

var MessageStore = []Message{}

func MessageHandler(m string, s string) (Message, error) {
	re := regexp.MustCompile(`^[a-zA-Z\s]*$`)

	if re.MatchString(m) {
		NewMessage := Message{
			MessageId:      0,
			MessageText:    m,
			MessageSender:  s,
			MessageTicket:  ticket.GenerateTicketNo(),
			Submitted:      time.Now(),
			Delivered:      time.Time{},
			DeliveredState: false,
		}

		db.InsertMessage(NewMessage.MessageText, NewMessage.MessageSender, NewMessage.MessageTicket, NewMessage.Submitted)

		log.Infof("added: %v, %d items in store", NewMessage, len(MessageStore))

		return NewMessage, nil
	} else {
		return Message{}, errors.New("input contains invalid characters")
	}
}

func AddToIMS(m Message) {
	MessageStore = append(MessageStore, m)
}

func CheckIMS(t string) Message {
	for _, m := range MessageStore {
		if m.MessageTicket == t {
			log.Infof("message found matching %s, returning message to caller", m.MessageTicket)
			return m
		}
	}
	return Message{0, "no message found", "", t, time.Time{}, time.Time{}, false}
}

func StringifyMessage(m Message) string {
	return fmt.Sprintf("Message: %s, from: %s (%s)", m.MessageText, m.MessageSender, m.MessageTicket)
}
