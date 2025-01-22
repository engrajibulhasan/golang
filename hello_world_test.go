package main

import "testing"

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHello(t *testing.T) {
	t.Run("saying hello to people===", func(t *testing.T) {
		output := message("", "Bangla")
		want := "আসসালামুয়ালাইকুম নওফা"
		assertCorrectMessage(t, output, want)
	})
	t.Run("saying hello world when an empty string is supplied===", func(t *testing.T) {
		output := message("", "")
		want := "Hello world"
		assertCorrectMessage(t, output, want)
	})
}
