package message

import (
	"fmt"
	"morseme/server/ticket"

	"github.com/labstack/gommon/log"
)

type Message struct {
	MessageText   string
	MessageSender string
	MessageTicket string
}

var MessageStore = []Message{}

func MessageHandler(m string, s string) string {
	NewMessage := Message{
		MessageText:   m,
		MessageSender: s,
		MessageTicket: ticket.GenerateTicketNo(),
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
	return Message{"no message found", "", ""}
}

func StringifyMessage(m Message) string {
	return fmt.Sprintf("Message: %s, from: %s (%s)", m.MessageText, m.MessageSender, m.MessageTicket)
}
