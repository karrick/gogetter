package gogetter

import "net/http"

// Getter is the interface implemented by an object that provides a Get method identical to
// http.Client's method, to allow composition of functionality.
type Getter interface {
	Get(url string) (*http.Response, error)
}
