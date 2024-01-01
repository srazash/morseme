package message

import (
	"fmt"
	"morseme/server/ticket"
	"time"

	"github.com/labstack/gommon/log"
)

type Message struct {
	MessageId      int
	MessageText    string
	MessageSender  string
	MessageTicket  string
	Submitted      time.Time
	Delivered      time.Time
	DeliveredState bool
}

var MessageStore = []Message{}

func MessageHandler(m string, s string) string {
	NewMessage := Message{
		MessageId:      len(MessageStore) + 1,
		MessageText:    m,
		MessageSender:  s,
		MessageTicket:  ticket.GenerateTicketNo(),
		Submitted:      time.Now(),
		Delivered:      time.Time{},
		DeliveredState: false,
	}

	MessageStore = append(MessageStore, NewMessage)

	log.Infof("added: %v, %d items in store", NewMessage, len(MessageStore))

	return fmt.Sprintf("%v\n", NewMessage)
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
	return Message{0, "no message found", "", "", time.Time{}, time.Time{}, false}
}

func StringifyMessage(m Message) string {
	return fmt.Sprintf("Message: %s, from: %s (%s)", m.MessageText, m.MessageSender, m.MessageTicket)
}

func MessageStats() (int, int, int) {
	total := 0
	undelivered := 0
	delivered := 0

	total = len(MessageStore)

	for _, v := range MessageStore {
		if !v.DeliveredState {
			undelivered += 1
		}
	}

	delivered = total - undelivered

	return total, undelivered, delivered
}
