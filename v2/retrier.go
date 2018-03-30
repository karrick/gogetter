package gogetter

import (
	"net/http"
	"time"
)

// Retrier will optionally retry queries that fail.
//
// WARNING: Retrier *should* wrap RoundRobin. That way retries don't go to the same dead
// server over and over.
type Retrier struct {
	// Getter is the object whose Get method is invoked to Get the results for a query.
	Getter Getter

	// RetryCount is number of query retries to be issued if query returns error. Leave 0 to
	// never retry query errors. But if you don't want to retry errors, It's best not to use a
	// Retrier...
	RetryCount int

	// RetryCallback is predicate function that tests whether query should be retried for a
	// given error. Leave nil to retry all errors.
	RetryCallback func(error) bool

	// RetryPause is the amount of time to wait before retrying the query with the underlying
	// Getter.
	RetryPause time.Duration
}

// Get attempts the specified query, and optionally retries a specified number of times, based on
// the results of calling the RetryCallback function.
func (r *Retrier) Get(url string) (response *http.Response, err error) {
	var attempts int
	for {
		response, err = r.Getter.Get(url)
		if err == nil {
			return
		}
		if attempts == r.RetryCount || r.RetryCallback == nil || r.RetryCallback(err) == false {
			return
		}
		attempts++
		if r.RetryPause > 0 {
			time.Sleep(r.RetryPause)
		}
	}
}
