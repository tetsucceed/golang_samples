package lesson6

import (
	"fmt"
	"log"
	"net/http"
)

func SimpleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	html_text := `<doctype html>
	<html>
	<head>
	<title>Hello %s</title>
	</head>
	<body>
		Hello %s
	</body>
	</html>`

	name, status := r.URL.Query()["name"]

	if !status || len(name[0]) < 1 {
		log.Println(fmt.Fprintf(w, html_text, "World!", "World!"))
		return
	}
	log.Println(fmt.Fprintf(w, html_text, name[0], name[0]))
}
