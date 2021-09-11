package shortlink

import (
	"fmt"
	"math/rand"
	"time"

	"amcosta.dev/encurtador-url/config"
)

type Generator struct {
	linkHash    string
	linkShorted string
}

func (s *Generator) New() {
	var hash string
	var randomIndex int

	rand.Seed(time.Now().UnixNano())
	for len(hash) < config.ShortLinkSize {
		randomIndex = rand.Intn(len(config.AvailableCharactersToRandom) - 1)
		hash += string(config.AvailableCharactersToRandom[randomIndex])
	}

	s.linkHash = hash
	s.linkShorted = fmt.Sprintf("%s/%s", config.Domain, hash)
}

func (sl *Generator) Hash() string {
	return sl.linkHash
}

func (sl *Generator) Url() string {
	return sl.linkShorted
}
