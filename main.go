package main

import (
	// "errors"
	"encoding/json"
	"fmt"

	// "io"
	"log"
	"net/http"
	// "os"
)

type Joke struct {
	IconUrl string
	Id      string
	URL     string
	Value   string
}

func main() {
	jokeText, err := jokeFetcher()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(jokeText)
}

func jokeFetcher() (string, error) {
	log.Println("Fetching Joke")
	jokeInstance := &Joke{}
	resp, err := http.Get("https://api.chucknorris.io/jokes/random")
	if err != nil {
		return "", err
	}

	err = json.NewDecoder(resp.Body).Decode(jokeInstance)
	return jokeInstance.Value, err
}
