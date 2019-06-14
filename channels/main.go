package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		"http://wikipedia.org",
	}

	c := make(chan string)

	// Request happens sequentially.
	// Have to wait for one to complete before starting next
	for _, link := range links {
		// checkLink(link)
		// go routine
		go checkLink(link, c)
	}
	// fmt.Println(<-c)

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			go checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "Might be down!")
		c <- link
		return
	}
	fmt.Println(link, "Link is up")
	c <- link

}
