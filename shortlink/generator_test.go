package shortlink_test

import (
	"strings"
	"testing"

	"amcosta.dev/encurtador-url/config"
	"amcosta.dev/encurtador-url/shortlink"
)

func TestCreateLink(t *testing.T) {
	link := &shortlink.ShortLink{}
	link.New()
	size := len(link.Hash())

	if size != config.ShortLinkSize {
		t.Errorf("the size of the hash must be %d, current size is %d", config.ShortLinkSize, size)
	}

	if !strings.Contains(link.Url(), config.Domain) {
		t.Errorf("the url doesn't contains the domain %s, current value is %s", config.Domain, link.Url())
	}
}
