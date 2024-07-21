package ascii_gen

import (
	"strings"
)

func FileSpliter(text string) [][]string {
	// we take text (banner without first line) as a string
	
	charsSlice := [][]string{} 
	charSlice := []string{}	

	for _, line := range strings.Split(text, "\n") {
		// We check if the line is empty (end of char) and if charSlice
		if line == "" && charSlice != nil {
			// Seconde we append the character in the characters holder and empty the char
			charsSlice = append(charsSlice, charSlice)
			charSlice = []string{}
		} else {
			// First we append all character lines in charSlice
			charSlice = append(charSlice, line)
		}
	}

	return charsSlice
}
