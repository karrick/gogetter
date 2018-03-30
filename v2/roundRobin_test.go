package gogetter_test

import (
	"testing"

	"github.com/karrick/gogetter"
)

func TestRoundRobinSingleGetter(t *testing.T) {
	getter := new(mockGetter)

	roundRobin := gogetter.NewRoundRobin([]gogetter.Getter{getter})

	_, _ = roundRobin.Get("")

	if actual, expected := getter.invokedCounter, 1; actual != expected {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}
}

func TestRoundRobinSingleGetterInvokesMultipleTimes(t *testing.T) {
	getter := new(mockGetter)

	roundRobin := gogetter.NewRoundRobin([]gogetter.Getter{getter})

	_, _ = roundRobin.Get("")
	_, _ = roundRobin.Get("")

	if actual, expected := getter.invokedCounter, 2; actual != expected {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}
}

func TestRoundRobinMultipleGetterInvokedMultipleTimes(t *testing.T) {
	getter1 := new(mockGetter)
	getter2 := new(mockGetter)
	getter3 := new(mockGetter)

	roundRobin := gogetter.NewRoundRobin([]gogetter.Getter{getter1, getter2, getter3})

	_, _ = roundRobin.Get("")
	_, _ = roundRobin.Get("")
	_, _ = roundRobin.Get("")
	_, _ = roundRobin.Get("")
	_, _ = roundRobin.Get("")

	if actual, expected := getter1.invokedCounter, 2; actual != expected {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}
	if actual, expected := getter2.invokedCounter, 2; actual != expected {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}
	if actual, expected := getter3.invokedCounter, 1; actual != expected {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}
}
