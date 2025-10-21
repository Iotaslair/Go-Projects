package main

import (
	"fmt"
	"net/http"
)

type result struct {
	url string
	status int
}

func main() {
	urls := []string {"https://httpco.de/200", "https://httpco.de/403", "https://httpco.de/404", }

	channel := make(chan result)

	for _, url := range urls {
		go makeCall(url, channel)
	}

	for range urls {
		result := <- channel
		fmt.Println(result.url, result.status)
	}
}

func makeCall(url string, channel chan result){
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	checkError(err)
	resp, err := client.Do(req)
	channel <- result{url, resp.StatusCode}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}