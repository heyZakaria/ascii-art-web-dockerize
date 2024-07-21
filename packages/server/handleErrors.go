package web

import (
	"fmt"
	"html/template"
	"net/http"
	data "web/packages/common_data"
)

type ErrorData struct {
	Msg      string
	Code     int
	DarkMode bool
}

func ErrorHandler(w http.ResponseWriter, msg string, code int) {
	errorMsg := ErrorData{msg, code, data.Result.DarkMode}

	w.WriteHeader(code)

	tmpl, err := template.ParseFiles("./pages/error.html")
	if err != nil {
		http.Error(w, msg, code)
		fmt.Println("Error When we try parse the file", err)
		return
	}

	err = tmpl.Execute(w, errorMsg)
	if err != nil {
		http.Error(w, msg, code)
		fmt.Println("Error When we try excute the file", err)
		return
	}
}
