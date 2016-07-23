package gogetter

import (
	"log"
	"net/http"
)

// Retrier will optionally retry queries that fail.
//
// WARNING: Retrier *should* wrap RoundRobin. That way retries don't go to the same dead
// server over and over.
type Retrier struct {
	Getter        Getter
	RetryCount    int
	RetryCallback func(error) bool // true means retry query based on specified error
}

// Get attempts the specified query, and optionally retries a specified number of times, based on
// the results of calling the RetryCallback function.
func (r *Retrier) Get(url string) (response *http.Response, err error) {
	log.Printf("Retrier.Get") // DEBUG
	// NOTE: condition is less than or equal to ensure it runs once _plus_ retry count
	for count := 0; count <= r.RetryCount; count++ {
		response, err = r.Getter.Get(url)
		if err == nil || (r.RetryCallback != nil && !r.RetryCallback(err)) {
			return
		}
		log.Print(err) // DEBUG
	}
	return
}
