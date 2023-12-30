package message

import (
	"fmt"
	"morseme/server/ticket"

	"github.com/labstack/gommon/log"
)

type Message struct {
	MessageId     int
	MessageText   string
	MessageSender string
	MessageTicket string
	Delivered     bool
}

var MessageStore = []Message{}

func MessageHandler(m string, s string) string {
	NewMessage := Message{
		MessageId:     len(MessageStore) + 1,
		MessageText:   m,
		MessageSender: s,
		MessageTicket: ticket.GenerateTicketNo(),
		Delivered:     false,
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
	return Message{0, "no message found", "", "", false}
}

func StringifyMessage(m Message) string {
	return fmt.Sprintf("Message: %s, from: %s (%s)", m.MessageText, m.MessageSender, m.MessageTicket)
}

func MessageStats() (int, int, int) {
	t := 0
	u := 0
	d := 0

	t = len(MessageStore)

	for _, v := range MessageStore {
		if !v.Delivered {
			u += 1
		}
	}

	d = t - u

	return t, u, d
}
