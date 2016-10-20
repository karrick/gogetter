package gogetter

import (
	"container/ring"
	"net/http"
)

// RoundRobin responds to the Get method by invoking all specified Getters in round-robin fashion.
type roundRobin struct {
	r *ring.Ring
}

// NewRoundRobin returns a Getter that sends successive queries to all the Getters its list.
func NewRoundRobin(getters []Getter) Getter {
	r := ring.New(len(getters))

	for _, getter := range getters {
		r = r.Next()
		r.Value = getter
	}

	return &roundRobin{r: r}
}

// Get responds to the Get method by invoking all specified Getters in round-robin fashion.
func (g *roundRobin) Get(url string) (*http.Response, error) {
	next := g.r.Next()
	g.r = next
	return next.Value.(Getter).Get(url)
}
