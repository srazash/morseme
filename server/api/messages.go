package api

import (
	"encoding/json"
	"log"
	"morseme/server/message"
	"time"
)

type Messages struct {
	Messages []struct {
		MessageId      int       `json:"message_id"`
		MessageText    string    `json:"message_text"`
		MessageSender  string    `json:"message_sender"`
		MessageTicket  string    `json:"message_ticket"`
		Submitted      time.Time `json:"submitted"`
		Delivered      time.Time `json:"delivered"`
		DeliveredState bool      `json:"delivered_state"`
	}
}

func MessagesJson(ms []message.Message) []byte {
	messagesJson, err := json.Marshal(ms)
	if err != nil {
		log.Fatalf("unable to marshal json: %v\n", err)
	}

	return messagesJson
}
