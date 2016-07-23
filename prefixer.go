package gogetter

import "net/http"

// Prefixer prepends the specified prefix to the URL before redirecting to the underlying Getter.
type Prefixer struct {
	Prefix string
	Getter Getter
}

// Get prepends the specified prefix to the URL before redirecting to the underlying Getter.
func (p *Prefixer) Get(url string) (*http.Response, error) {
	return p.Getter.Get(p.Prefix + url)
}
