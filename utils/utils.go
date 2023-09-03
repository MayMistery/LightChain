package utils

import (
	"net/http"
	"strings"
)

func IndexHtml(w http.ResponseWriter) {
	var html strings.Builder

	html.WriteString("<html><head><title>Home Page</title></head><body>")

	html.WriteString(`<div style="font-family:cursive;color:red;font-size:48px;">`)
	html.WriteString("Red Server")
	html.WriteString(`</div>`)
	html.WriteString("</body></html>")

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(html.String()))
}
