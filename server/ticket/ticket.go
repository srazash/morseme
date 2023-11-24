package ticket

import (
	"math/rand"
	"strings"
)

const charset = "1234567890abcdefghijklmnopqrstuvwxyz"

func GenerateTicketNo() string {
	ticket := strings.Builder{}
	ticket.Grow(9)

	for i := 0; i < 9; i++ {
		if i == 4 {
			ticket.WriteByte('-')
		} else {
			ticket.WriteByte(charset[rand.Intn(len(charset))])
		}
	}

	return strings.ToUpper(ticket.String())
}
