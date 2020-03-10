package lesson8

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func ParallelScan(sites []string) {
	start := time.Now()
	ch := make(chan string)
	for _, url := range sites {
		go fetch(url, ch)
	}
	for range sites {
		fmt.Print(<-ch) // receive from channel
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	// тут можно собирать урлы в теле сайта и ползти дальше
	nbytes, err := io.Copy(ioutil.Discard, resp.Body) // to devNull
	// условно
	// urls := parse(resp.Body)
	// fetch(urls, ch)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	ch <- fmt.Sprintf("%.2fs %7d %s\n", time.Since(start).Seconds(), nbytes, url)
}
