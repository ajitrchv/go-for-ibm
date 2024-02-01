package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://github.com",
		"http://fb.com",
		"http://golang.org",
		"http://youtube.com",
		"http://amazon.com",
		"http://stackoverflow.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		out := fmt.Sprintf("%s might be down", link)
		c <- out
		return
	}
	out := fmt.Sprintf("%s is up!!", link)
	c <- out
}
