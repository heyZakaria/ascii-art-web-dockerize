package web

import (
	"net/http"

	ascii_gen "web/packages/ascii_gen/src"
	data "web/packages/common_data"
)

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		text := r.FormValue("text")
		banner := r.FormValue("banner")
		path := "./packages/ascii_gen/banners/" + banner + ".txt"
		if text == "" || banner == "" || !ascii_gen.CheckBanner(path) || len(text) > 2000 {
			ErrorHandler(w, "Bad Request", http.StatusBadRequest)
			return
		}

		data.Result.Text = text
		data.Result.Banner = banner

		http.Redirect(w, r, "/", http.StatusSeeOther)

	} else {
		ErrorHandler(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
