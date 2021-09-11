// Essa documentação é do repository
package repository

import (
	"fmt"

	"amcosta.dev/encurtador-url/errors"
	"amcosta.dev/encurtador-url/models"
)

var links []models.Url

type InMemory struct {
}

func (s *InMemory) FindByOriginalLink(link string) (models.Url, error) {
	for _, url := range links {
		if url.OriginalLink == link {
			return url, nil
		}
	}

	return models.Url{}, fmt.Errorf("the url %s doesn't exist", link)
}

func (s *InMemory) FindByShortedLink(link string) (models.Url, error) {
	for _, url := range links {
		if url.ShortedLink == link {
			return url, nil
		}
	}

	return models.Url{}, fmt.Errorf("the url %s doesn't exist", link)
}

func (s *InMemory) FindAll() []models.Url {
	return links
}

func (s *InMemory) Persist(url models.Url) (models.Url, error) {
	if url.Id == "" {
		return models.Url{}, fmt.Errorf("url must contain a Id, use the url hash for a valid Id")
	}

	links = append(links, url)
	return url, nil
}

func (s *InMemory) FindByHash(hash string) (models.Url, error) {
	for _, url := range links {
		if url.Id == hash {
			return url, nil
		}
	}

	return models.Url{}, errors.ErrModelNotFoundByHash
}
