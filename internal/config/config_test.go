package config 

import "testing"

func TestWorkers(t *testing.T) {
	got := Workers()
	expected := 4

	if got != expected {
		t.Errorf("Expected number of workers = %q, Got: %q", expected, got)
	}
}
