package gogetter_test

import (
	"errors"
	"net/http"
)

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
