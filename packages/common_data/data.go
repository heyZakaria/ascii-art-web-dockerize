package data

import (
	"strings"
)

type Data struct {
	Text            string
	Banner          string
	Ascii_art       string
	Error           string
	NotAllowedChars []string
	DarkMode        bool
}

var Result = Data{}

func HandleErrorData(err []string) string {
	suffix := ""
	chars := MakeItUnique(err)
	if len(chars) == 1 {
		suffix = "This character doesn't exist in the ASCII table."
	} else if len(chars) > 1 {
		suffix = "These characters don't exist in the ASCII table."
	} else {
		return ""
	}

	handledError := strings.Join(chars, ", ")

	return handledError + ": " + suffix
}

func CleanData(data *Data) {
	data.Text = ""
	data.Banner = ""
	data.Ascii_art = ""
	data.Error = ""
	data.NotAllowedChars = nil
}
