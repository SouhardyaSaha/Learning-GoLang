package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	channel := make(chan string)

	for _, link := range links {
		go checkLink(link, channel)
	}

	for link := range channel {
		go func(linkInsideFunction string) {
			time.Sleep(5 * time.Second)
			checkLink(linkInsideFunction, channel)
		}(link)

	}

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-channel)
	// }
}

func checkLink(link string, channel chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "is down!\n", err)
		channel <- link
		return
	}
	fmt.Println(link, "is up!")
	channel <- link
}
