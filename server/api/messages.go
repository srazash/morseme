package api

import (
	"encoding/json"
	"log"
	"morseme/server/db"
)

func MessagesJson(m []db.Message) []byte {
	messagesJson, err := json.Marshal(m)
	if err != nil {
		log.Fatalf("unable to marshal json: %v\n", err)
	}

	return messagesJson
}

func MessageJson(m db.Message) []byte {
	messageJson, err := json.Marshal(m)
	if err != nil {
		log.Fatalf("unable to marshal json: %v\n", err)
	}

	return messageJson
}
