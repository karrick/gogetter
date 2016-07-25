package gogetter

import "net/http"

// Retrier will optionally retry queries that fail.
//
// WARNING: Retrier *should* wrap RoundRobin. That way retries don't go to the same dead
// server over and over.
type Retrier struct {
	Getter        Getter
	RetryCount    int              // RetryCount is number of query retries to be issued if query returns error. Leave 0 to never retry query errors. But if you don't want to retry errors, It's best not to use a Retrier...
	RetryCallback func(error) bool // RetryCallback is predicate function that tests whether query should be retried for a given error. Leave nil to retry all errors.
}

// Get attempts the specified query, and optionally retries a specified number of times, based on
// the results of calling the RetryCallback function.
func (r *Retrier) Get(url string) (response *http.Response, err error) {
	// NOTE: condition is less than or equal to ensure it runs once _plus_ retry count
	for count := 0; count <= r.RetryCount; count++ {
		response, err = r.Getter.Get(url)
		if err == nil || (r.RetryCallback != nil && !r.RetryCallback(err)) {
			return
		}
	}
	return
}
