package gogetter_test

import (
	"testing"

	"github.com/karrick/gogetter"
)

func TestRetrierZeroRetriesMakesOnlyOneAttempt(t *testing.T) {
	var retryCalbackInvocationCounter int
	getter := &mockGetter{failuresRemaining: 10}
	retrier := &gogetter.Retrier{
		Getter:     getter,
		RetryCount: 0, // want max of 1 attempt
		RetryCallback: func(_ error) bool {
			retryCalbackInvocationCounter++
			return true
		},
	}

	_, _ = retrier.Get("")

	if actual, expected := getter.invokedCounter, 1; actual != expected {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}
	if actual, expected := retryCalbackInvocationCounter, 0; actual != expected {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}
}

func TestRetrierOnlyInvokesCallbackWhenMoreRetriesLeft(t *testing.T) {
	var retryCalbackInvocationCounter int
	getter := &mockGetter{failuresRemaining: 10}
	retrier := &gogetter.Retrier{
		Getter:     getter,
		RetryCount: 2, // want max of 3 attempts
		RetryCallback: func(_ error) bool {
			retryCalbackInvocationCounter++
			return true
		},
	}

	_, _ = retrier.Get("")

	if actual, expected := getter.invokedCounter, 3; actual != expected {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}
	if actual, expected := retryCalbackInvocationCounter, 2; actual != expected {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}
}

func TestRetrierOnlyRetriesWhenCallbackReturnsTrue(t *testing.T) {
	var retryCalbackInvocationCounter int
	getter := &mockGetter{failuresRemaining: 10}
	retrier := &gogetter.Retrier{
		Getter:     getter,
		RetryCount: 2, // want max of 3 attempts
		RetryCallback: func(_ error) bool {
			retryCalbackInvocationCounter++
			return false
		},
	}

	_, _ = retrier.Get("")

	if actual, expected := getter.invokedCounter, 1; actual != expected {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}
	if actual, expected := retryCalbackInvocationCounter, 1; actual != expected {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}
}

func TestRetrierEventualSuccess(t *testing.T) {
	var retryCalbackInvocationCounter int
	getter := &mockGetter{failuresRemaining: 2}
	retrier := &gogetter.Retrier{
		Getter:     getter,
		RetryCount: 4, // want max of 5 attempts
		RetryCallback: func(_ error) bool {
			retryCalbackInvocationCounter++
			return true
		},
	}

	_, _ = retrier.Get("")

	if actual, expected := getter.invokedCounter, 3; actual != expected {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}
	if actual, expected := retryCalbackInvocationCounter, 2; actual != expected {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}
}
