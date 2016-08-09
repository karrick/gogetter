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

	// log.Printf("servers: %v\n", getters)
	for i := 0; i < len(getters); i++ {
		r.Next().Value = getters[i]
	}

	return &roundRobin{r: r}
}

// Get responds to the Get method by invoking all specified Getters in round-robin fashion.
func (g *roundRobin) Get(url string) (*http.Response, error) {
	return g.r.Next().Value.(Getter).Get(url)
}
