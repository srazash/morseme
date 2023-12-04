package morsecode

import "strings"

func Encode(input string) string {

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

	return output
}
