package gogetter

import (
	"net/http"
	"sync"
)

// RoundRobin responds to the Get method by invoking all specified Getters in round-robin fashion.
type RoundRobin struct {
	Getters   []Getter
	indexLock sync.Mutex
	index     uint64
}

// Get responds to the Get method by invoking all specified Getters in round-robin fashion.
func (g *RoundRobin) Get(url string) (*http.Response, error) {
	g.indexLock.Lock()
	index := g.index
	g.index++
	if g.index == uint64(len(g.Getters)) {
		g.index = 0
	}
	g.indexLock.Unlock() // NOTE: indexLock is not unlocked using defer because the Get invocation below can take quite a bit of time.
	return g.Getters[index].Get(url)
}
