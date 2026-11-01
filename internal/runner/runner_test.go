package runner

import "testing"

func TestBanner(t *testing.T) {
	got := Banner()
	want := `Task Runner starting...
Project version: 1.0`

	if got != want {
		t.Errorf("Banner() = %q, want = %q", got, want)
	}
}

func TestName(t *testing.T) {
	got := Name()
	want := "taskrunner"

	if got != want {
		t.Errorf("Want = %q, got = %q", want, got)
	}
}
