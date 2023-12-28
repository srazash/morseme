package message

import (
	"fmt"
	"morseme/server/ims"
	"morseme/server/ticket"

	"github.com/labstack/gommon/log"
)

func MessageHandler(m string, s string) string {
	NewMessage := ims.Message{
		MessageText:   m,
		MessageSender: s,
		MessageTicket: ticket.GenerateTicketNo(),
	}

	ims.MessageStore = append(ims.MessageStore, NewMessage)

	log.Infof("added: %v, %d items in store", NewMessage, len(ims.MessageStore))

	return fmt.Sprintf("%v\n", NewMessage)
}
