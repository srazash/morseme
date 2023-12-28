package ims

type Message struct {
	MessageText   string
	MessageSender string
	MessageTicket string
}

var MessageStore = []Message{}

func AddToIMS(m Message) {
	MessageStore = append(MessageStore, m)
}

func CheckIMS(t string) Message {
	for _, m := range MessageStore {
		if m.MessageTicket == t {
			return m
		}
	}
	return Message{"", "", ""}
}
