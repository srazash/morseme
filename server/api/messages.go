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

func LastMessageJson(ms []message.Message) []byte {
	i := len(ms) - 1

	messagesJson, err := json.Marshal(ms[i])
	if err != nil {
		log.Fatalf("unable to marshal json: %v\n", err)
	}

	return messagesJson
}

func FirstUndeliveredMessageJson(ms []message.Message) []byte {
	i := len(ms) - 1

	for i >= 0 {
		if !ms[i].DeliveredState {
			i--
		} else {
			i++
			break
		}
	}

	messagesJson, err := json.Marshal(ms[i])
	if err != nil {
		log.Fatalf("unable to marshal json: %v\n", err)
	}

	return messagesJson
}
