package reflector

import "testing"

func TestWalk(t *testing.T) {
	var got []string
	expected := "Susan"
	sutStruct := struct {
		field1 string
	}{expected}

	want := 1
	Walk(sutStruct, func(s string) {
		got = append(got, s)
	})

	if len(got) != want {
		t.Fatalf("Got %d calls instead of %d", len(got), want)
	}

	if got[0] != expected {
		t.Errorf("got %q, want %q", got[0], expected)
	}
}
