package main

import (
	"fmt"
	"lesson6"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Starting lesson6...")

	// draw two crossed lines
	lesson6.Draw()

	// prepare for .... magic
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	if port == 0 || port < 1024 {
		log.Fatal("$PORT is lower than 1024")
	}

	real_port := fmt.Sprintf(":%d", port)

	// TODO: rewrite it with gin
	http.HandleFunc("/hello", lesson6.SimpleHandler)
	log.Fatal(http.ListenAndServe(real_port, nil))

}
