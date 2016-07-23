package gogetter

import (
	"log"
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
	log.Printf("RoundRobin.Get") // DEBUG
	g.indexLock.Lock()
	index := g.index
	g.index++
	if g.index == uint64(len(g.Getters)) {
		g.index = 0
	}
	g.indexLock.Unlock()
	return g.Getters[index].Get(url)
}
