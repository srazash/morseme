package api

import (
	"encoding/json"
	"log"
	"morseme/server/message"
)

func MessagesJson(ms []message.Message) []byte {
	messagesJson, err := json.Marshal(ms)
	if err != nil {
		log.Fatalf("unable to marshal json: %v\n", err)
	}

	return messagesJson
}
