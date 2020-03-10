package lesson8

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func ConcurrentScan(sites []string) {
	start := time.Now()
	for _, url := range sites {
		fetch2(url)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch2(url string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	//bytes, err := ioutil.ReadAll(resp.Body)
	nbytes, err := io.Copy(ioutil.Discard, resp.Body) // to devNull
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.2fs %7d %s\n", time.Since(start).Seconds(), nbytes, url)
}
