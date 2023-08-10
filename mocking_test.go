package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	CountDown(buffer)
	got := buffer.String()
	want := "3"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestCountDown(t *testing.T) {
	buffer := &bytes.Buffer{}

	CountDown(buffer)

	got := buffer.String()
	want := `3
2
1
Go!`
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
