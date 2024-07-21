package ascii_gen

import (
	data "web/packages/common_data"
)

func Writer(text string, banner [][]string) (string, error) {
	var textSlice [][]string

	var asciiText string

	for i, char := range text {

		if char == '\n' {
			if textSlice != nil {
				asciiText += lineWriter(textSlice)
				textSlice = nil
			}
			if (i+2 < len(text) && text[i+1:i+3] == "\r\n") || i+1 >= len(text) {
				asciiText += "\r\n"
			}
			continue
		} else if char == '\r' {
			continue
		} else if !IsChar(char) {
			data.Result.NotAllowedChars = append(data.Result.NotAllowedChars, string(char))
			continue
		}

		textSlice = append(textSlice, banner[int(char-32)])
	}

	if textSlice != nil {
		asciiText += lineWriter(textSlice)
	}

	return asciiText, nil
}

func lineWriter(text [][]string) string {
	ascii_text := ""

	for i := 0; i < 8 && text != nil; i++ {
		for index := range text {
			ascii_text += text[index][i]
		}
		ascii_text += "\n"
	}

	return ascii_text
}

func IsChar(char rune) bool {
	return char >= ' ' && char <= '~'
}
