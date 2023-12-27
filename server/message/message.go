package message

import (
	"fmt"
	"morseme/server/ticket"
)

type Message struct {
	MessageText   string
	MessageSender string
	MessageTicket string
}

func MessageHandler(m string, s string) string {
	NewMessage := Message{
		MessageText:   m,
		MessageSender: s,
		MessageTicket: ticket.GenerateTicketNo(),
	}

	return fmt.Sprintf("%v\n", NewMessage)
}
