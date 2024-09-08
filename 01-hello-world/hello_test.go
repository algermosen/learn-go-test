package helloworld

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {

		got := Hello("Alejandro", "")
		want := "Hello, Alejandro"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, world' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Juan", "Spanish")
		want := "Hola, Juan"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Julien", "French")
		want := "Bonjour, Julien"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Italian", func(t *testing.T) {
		got := Hello("Julio", "Italian")
		want := "Ciao, Julio"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
