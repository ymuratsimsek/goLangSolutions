package main

// Fetchall fetches URLs in parallel and reports their times and sizes.

import (
	"fmt"
	"io"
	"net/http"
	net "net/url"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) //start go routine
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) //receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) //send to channel
		return
	}

	u, _ := net.Parse(url)               // getting hostname from url
	file, err := os.Create(u.Hostname()) // creating file by using hostname as a filename
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(file, resp.Body)
	resp.Body.Close() //dont leak resources
	if err != nil {
		ch <- fmt.Sprintf("While reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)

}
