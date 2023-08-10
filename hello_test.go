package main

import (
	"strings"
	"testing"
)

//func TestHello(t *testing.T) {
//	got := Hello("Hello, world")
//	want := "Hello, world"
//
//	if got != want {
//		t.Errorf("got %q want %q", got, want)
//	}
//}

func TestHello(t *testing.T) {
	t.Run("saying hello to perhaps", func(t *testing.T) {
		got := Hello("Hello")
		want := "Hello"
		if !strings.EqualFold(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Hello", func(t *testing.T) {
		got := Hello("Hello")
		want := "Hello"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func TestRepeat(t *testing.T) {
	requested := Repeat("a")
	expected := "aaaaa"
	if requested != expected {
		t.Errorf("expected %q but got %q", expected, requested)
	}
}

func Repeat(s string) interface{} {
	var re string
	for i := 0; i < 5; i++ {
		re += s
	}
	return re
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

func Hello(s string) string {
	return s
}
