package gogetter

import (
	"log"
	"net/http"
)

// Prefixer prepends the specified prefix to the URL before redirecting to the underlying Getter.
type Prefixer struct {
	Prefix string
	Getter Getter
}

// Get prepends the specified prefix to the URL before redirecting to the underlying Getter.
func (p *Prefixer) Get(url string) (*http.Response, error) {
	log.Printf("Prefixer.Get") // DEBUG
	return p.Getter.Get(p.Prefix + url)
}
