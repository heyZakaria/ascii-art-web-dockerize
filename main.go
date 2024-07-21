package main

import (
	"fmt"
	"net/http"

	web "web/packages/server"
)

func main() {
	fmt.Println("http://localhost:4050")
	staticFolder := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", staticFolder))
	http.HandleFunc("/", web.HomeHandler)
	http.HandleFunc("/dark-mode", web.DarkMoodHandler)
	http.HandleFunc("/ascii-art", web.AsciiArtHandler)
	http.ListenAndServe(":4050", nil)
}
