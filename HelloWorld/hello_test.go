package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris!"
		assertMessage(t, got, want)
	})
	t.Run("Saying hello, but with empty string defaults to \"Hello World\"", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		assertMessage(t, got, want)
	})
	t.Run("In spanish", func(t *testing.T) {
		got := Hello("Mia", "spanish")
		want := "Hola, Mia!"
		assertMessage(t, got, want)
	})
	t.Run("In french", func(t *testing.T) {
		got := Hello("Pierre", "french")
		want := "Bounjour, Pierre!"
		assertMessage(t, got, want)
	})

}

func assertMessage(t testing.TB, got, want string) {
	if got != want {
		t.Helper() //If test fails inside this method, the line that calls it will fail, not this helper method
		t.Errorf("got %q, want %q", got, want)
	}
}
