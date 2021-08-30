package repository

import (
	"fmt"

	"amcosta.dev/encurtador-url/models"
	"github.com/google/uuid"
)

var links []models.UrlModel

func FindByOriginalLink(link string) (models.UrlModel, error) {
	for _, url := range links {
		if url.OriginalLink == link {
			return url, nil
		}
	}

	return models.UrlModel{}, fmt.Errorf("the url %s doesn't exist", link)
}

func FindByShortedLink(link string) (models.UrlModel, error) {
	for _, url := range links {
		if url.ShortedLink == link {
			return url, nil
		}
	}

	return models.UrlModel{}, fmt.Errorf("the url %s doesn't exist", link)
}

func CreateUrl(originalLink string, shortedLink string) models.UrlModel {
	url := models.UrlModel{
		Id:           uuid.NewString(),
		OriginalLink: originalLink,
		ShortedLink:  shortedLink,
	}

	links = append(links, url)
	return url
}

func FindAll() []models.UrlModel {
	return links
}
