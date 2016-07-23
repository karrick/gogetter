package gogetter

import (
	"fmt"
	"math/rand"
	"net/http"
)

// Failer is used for testing by causing periodic Get failures.
//
// WARNING: Don't use this in production code.
type Failer struct {
	Getter    Getter
	Frequency float64
}

// Get periodically responds to the Get method, and periodically fails.
func (f *Failer) Get(url string) (*http.Response, error) {
	if rand.Float64() < f.Frequency {
		return nil, fmt.Errorf("random error")
	}
	return f.Getter.Get(url)
}
