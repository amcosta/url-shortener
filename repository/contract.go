package repository

import "amcosta.dev/encurtador-url/models"

type UrlRepository interface {
	FindByOriginalLink(link string) (models.Url, error)
	Persist(url models.Url) (models.Url, error)
	FindByHash(hash string) (models.Url, error)
}
