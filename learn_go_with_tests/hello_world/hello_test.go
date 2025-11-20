package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to a person", func(t *testing.T) {
		got := Hello("William", "")
		want := "Hello William"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say Hello World when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Japanese", func(t *testing.T) {
		got := Hello("Tanaka", "Japanese")
		want := "Konnichiwa Tanaka"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Baguete", "French")
		want := "Bonjour Baguete"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
