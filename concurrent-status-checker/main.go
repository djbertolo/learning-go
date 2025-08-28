package main

import (
	"fmt"
	"net/http"
)

var websites []string = []string{
	"https://www.google.com",
	"https://www.github.com",
	"http://non-existent-site.dev",
}

func main() {

	websiteStatus := make(chan string)
	for _, v := range websites {
		go func(website string) {
			_, err := http.Get(website)
			if err != nil {
				websiteStatus <- website + " is DOWN!"
			} else {
				websiteStatus <- website + " is UP!"
			}
		}(v)
	}

	for range websites {
		fmt.Println(<-websiteStatus)
	}

	close(websiteStatus)

}
