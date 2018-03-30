package gogetter

import "net/http"

// Getter is very small abstraction of any type that provides a Get method with the same signature
// of http.Client's Get method.
type Getter interface {
	Get(url string) (*http.Response, error)
}
