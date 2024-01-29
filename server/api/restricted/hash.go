package restricted

import (
	"encoding/hex"

	"github.com/jzelinskie/whirlpool"
)

func HashString(in string) string {
	w := whirlpool.New()
	input := []byte(in)
	w.Write(input)
	output := hex.EncodeToString(w.Sum(nil))
	return output
}
