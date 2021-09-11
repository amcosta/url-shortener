package models

import "github.com/google/uuid"

type Url struct {
	Id           string
	OriginalLink string
	ShortedLink  string
}

func (url *Url) Init() {
	url.Id = uuid.NewString()
}
