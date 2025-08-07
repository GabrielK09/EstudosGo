package http

import (
	"fmt"
	"html"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

}
