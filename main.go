package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

const availableCharacters = "ABCDEFGHIJKLMNOPQRSTUVXWYZabcdefghijklmnopqrstuvwxyz1234567890-_"
const shortLinkSize = 8
const host = "https://short.link/"

type Url struct {
	Id           string `json:"id"`
	OriginalLink string `json:"original_link"`
	ShortedLink  string `json:"shorted_link"`
}

func main() {
	link := "https://www.google.com"

	url := Url{uuid.NewString(), link, fmt.Sprintf("%s%s", host, generateShortedLink())}

	jsonData, err := json.Marshal(url)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonData))
}

func generateShortedLink() string {
	var shortedLink string
	var randomIndex int

	rand.Seed(time.Now().UnixNano())
	for len(shortedLink) < shortLinkSize {
		randomIndex = rand.Intn(len(availableCharacters) - 1)
		shortedLink += string(availableCharacters[randomIndex])
	}

	return shortedLink
}
