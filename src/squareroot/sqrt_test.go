package main

import "testing"

func TestSqrtNTimes(t *testing.T) {
	actual := sqrtNTimes(256, 3)
	expected := 2.0
	if actual != expected {
		t.Errorf("Expected [%v], got [%v]", expected, actual)
	}
}