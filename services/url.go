package services

import (
	"amcosta.dev/encurtador-url/models"
	"amcosta.dev/encurtador-url/repository"
	"amcosta.dev/encurtador-url/shortlink"
)

type service struct {
	repository repository.UrlRepository
	shortLink  shortlink.Generator
}

func NewUrlService(repository repository.UrlRepository, hash shortlink.Generator) *service {
	return &service{
		repository: repository,
		shortLink:  hash,
	}
}

// Create a new url
func (s *service) Create(targetUrl string) models.Url {
	s.shortLink.New()
	_, err := s.repository.FindByHash(s.shortLink.Hash())
	if err == nil {
		return s.Create(targetUrl)
	}

	model := models.Url{Id: s.shortLink.Hash(), OriginalLink: targetUrl, ShortedLink: s.shortLink.Url()}
	_, err = s.repository.Persist(model)
	if err != nil {
		panic(err)
	}

	return model
}
