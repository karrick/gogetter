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
	Getter    Getter  // Getter's Get method is invoked when the Failer instance doesn't fail.
	Frequency float64 // Frequency specifies how often a failure should occur. Set to 0.0 to never fail and 1.0 to always fail.
}

// Get periodically responds to the Get method, and periodically fails.
func (f *Failer) Get(url string) (*http.Response, error) {
	if rand.Float64() < f.Frequency {
		return nil, fmt.Errorf("random error")
	}
	return f.Getter.Get(url)
}
