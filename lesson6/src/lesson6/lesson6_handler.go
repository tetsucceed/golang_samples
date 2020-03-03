package lesson6

import (
	"fmt"
	"log"
	"net/http"
)

func SimpleHandler(w http.ResponseWriter, r *http.Request) {
	name, status := r.URL.Query()["name"]

	if !status || len(name[0]) < 1 {
		log.Println(fmt.Fprintf(w, "name parameter is not set"))
		return
	}
	log.Println(fmt.Fprintf(w, "Hi there, I love %s!", name[0]))
}
