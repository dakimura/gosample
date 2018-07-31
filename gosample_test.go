package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Errorf("1+2 should be 3. actual: %v", Add(1, 2))
	} else {
		t.Log("passed my first test in golang.")
	}
}
