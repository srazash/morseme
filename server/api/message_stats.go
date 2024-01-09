package api

import (
	"encoding/json"
	"log"
)

type MessageStatsTotal struct {
	Total int `json:"total"`
}

type MessageStatsUndelivered struct {
	Undelivered int `json:"undelievered"`
}

type MessageStatsDelivered struct {
	Delivered int `json:"delivered"`
}

type MessageStats struct {
	Total       int `json:"total"`
	Undelivered int `json:"undelievered"`
	Delivered   int `json:"delivered"`
}

func MessageStatsTotalJson(total int) []byte {
	statsStruct := MessageStatsTotal{
		Total: total,
	}

	statsJson, err := json.Marshal(statsStruct)
	if err != nil {
		log.Fatalf("unable to marshal to json: %v\n", err)
	}

	return statsJson
}

func MessageStatsUndeliveredJson(undelivered int) []byte {
	statsStruct := MessageStatsUndelivered{
		Undelivered: undelivered,
	}

	statsJson, err := json.Marshal(statsStruct)
	if err != nil {
		log.Fatalf("unable to marshal to json: %v\n", err)
	}

	return statsJson
}

func MessageStatsDeliveredJson(delivered int) []byte {
	statsStruct := MessageStatsDelivered{
		Delivered: delivered,
	}

	statsJson, err := json.Marshal(statsStruct)
	if err != nil {
		log.Fatalf("unable to marshal to json: %v\n", err)
	}

	return statsJson
}

func MessageStatsJson(total int, undelivered int, delivered int) []byte {
	statsStruct := MessageStats{
		Total:       total,
		Undelivered: undelivered,
		Delivered:   delivered,
	}

	statsJson, err := json.Marshal(statsStruct)
	if err != nil {
		log.Fatalf("unable to marshal to json: %v\n", err)
	}

	return statsJson
}
