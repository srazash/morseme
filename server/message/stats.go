package message

func MessageStatsTotal() int {
	return len(MessageStore)
}

func MessageStatsUndelivered() int {
	undelivered := 0

	for _, v := range MessageStore {
		if !v.DeliveredState {
			undelivered += 1
		}
	}

	return undelivered
}

func MessageStatsDelivered() int {
	total := MessageStatsTotal()
	undelivered := MessageStatsUndelivered()

	return total - undelivered
}

func MessageStats() (int, int, int) {
	total := MessageStatsTotal()
	undelivered := MessageStatsUndelivered()
	delivered := MessageStatsDelivered()

	return total, undelivered, delivered
}
