package gogetter_test

import (
	"errors"
	"net/http"
)

// mockGetter allows mocking out the response of invoking a Getter by another Getter. If initialized
// with `failuresRemaining` greater than 0, it will fail that many times before succeeding. Note
// even after success the pointer to the `http.Response` remains nil.
type mockGetter struct {
	failuresRemaining int
	invokedCounter    int
}

func (g *mockGetter) Get(url string) (*http.Response, error) {
	g.invokedCounter++

	g.failuresRemaining--
	if g.failuresRemaining >= 0 {
		return nil, errors.New("intentional error")
	}
	return nil, nil
}
