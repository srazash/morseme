package message

import (
	"errors"
	"morseme/server/db"
	"morseme/server/ticket"
	"regexp"
	"time"
)

func MessageHandler(m string, s string) (db.Message, error) {
	re := regexp.MustCompile(`^[a-zA-Z\s]*$`)

	if re.MatchString(m) {
		NewMessage := db.Message{
			Id:             0,
			Message:        m,
			Sender:         s,
			Ticket:         ticket.GenerateTicketNo(),
			Submitted:      time.Now(),
			Delivered:      time.Time{},
			DeliveredState: false,
		}

		db.InsertMessage(NewMessage.Message, NewMessage.Sender, NewMessage.Ticket, NewMessage.Submitted)

		return NewMessage, nil
	} else {
		return db.Message{}, errors.New("input contains invalid characters")
	}
}
