package morsecode

import (
	"errors"
	"regexp"
	"strings"
)

func Encode(input string) (string, error) {

	re := regexp.MustCompile(`^[a-zA-Z\s]*$`)

	runeToMorseStr := map[rune]string{
		'a': ".-",
		'b': "-...",
		'c': "-.-.",
		'd': "-..",
		'e': ".",
		'f': "..-.",
		'g': "--.",
		'h': "....",
		'i': "..",
		'j': ".---",
		'k': "-.-",
		'l': ".-..",
		'm': "--",
		'n': "-.",
		'o': "---",
		'p': ".--.",
		'q': "--.-",
		'r': ".-.",
		's': "...",
		't': "-",
		'u': "..-",
		'v': "...-",
		'w': ".--",
		'x': "-..-",
		'y': "-.--",
		'z': "--..",
	}

	if !re.MatchString(input) {
		return "", errors.New("input contains invalid characters")
	}

	lowerInput := strings.ToLower(input)
	output := ""

	for i, r := range lowerInput {
		switch r {
		case ' ':
			continue
		default:
			output += runeToMorseStr[r]
		}

		if i < len(input)-1 {
			output += " "
		}
	}

	return output, nil
}
