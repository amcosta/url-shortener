package shortlink

import (
	"fmt"
	"math/rand"
	"time"
)

var availableCharacters = "ABCDEFGHIJKLMNOPQRSTUVXWYZabcdefghijklmnopqrstuvwxyz1234567890-_"
var shortLinkSize = 8
var domain = "https://short.link"

type ShortLink struct {
	LinkDomain  string
	LinkHash    string
	LinkShorted string
}

func New() ShortLink {
	var hash string
	var randomIndex int

	rand.Seed(time.Now().UnixNano())
	for len(hash) < shortLinkSize {
		randomIndex = rand.Intn(len(availableCharacters) - 1)
		hash += string(availableCharacters[randomIndex])
	}

	return ShortLink{
		LinkDomain:  domain,
		LinkHash:    hash,
		LinkShorted: fmt.Sprintf("%s/%s", domain, hash),
	}
}
