package gogetter

import "net/http"

// Prefixer prepends the specified prefix to the URL before redirecting to the underlying Getter.
type Prefixer struct {
	// The object whose Get method is invoked to get results for a query.
	Getter Getter

	// Prefix for each query, commonly the hostname.
	Prefix string
}

// Get prepends the specified prefix to the URL before redirecting to the underlying Getter.
func (p *Prefixer) Get(url string) (*http.Response, error) {
	return p.Getter.Get(p.Prefix + url)
}
