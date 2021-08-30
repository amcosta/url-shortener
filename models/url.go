package models

import "github.com/google/uuid"

type UrlModel struct {
	Id           string
	OriginalLink string
	ShortedLink  string
}

func (url *UrlModel) Init() {
	url.Id = uuid.NewString()
}
