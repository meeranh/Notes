package main

import (
	"fmt"
	"net/http"
	"sync"
)

var waitGroup sync.WaitGroup

func main() {
	webSites := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.youtube.com",
	}

	for _, url := range webSites {
		waitGroup.Add(1)
		go getStatusCode(url)
	}
	waitGroup.Wait()
}

func getStatusCode(url string) {
	defer waitGroup.Done()
	res, err := http.Get(url)
	errCheck(err)
	fmt.Printf("Returned code %d for URL %s\n", res.StatusCode, url)
}

func errCheck(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
