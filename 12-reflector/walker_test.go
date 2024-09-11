package reflector

import (
	"slices"
	"testing"
)

// func TestWalk(t *testing.T) {
// 	var got []string
// 	expected := "Susan"
// 	sutStruct := struct {
// 		field1 string
// 	}{expected}

// 	want := 1
// 	Walk(sutStruct, func(s string) {
// 		got = append(got, s)
// 	})

// 	if len(got) != want {
// 		t.Fatalf("Got %d calls instead of %d", len(got), want)
// 	}

// 	if got[0] != expected {
// 		t.Errorf("got %q, want %q", got[0], expected)
// 	}
// }

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one field string",
			struct{ Name string }{"Susan"},
			[]string{"Susan"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			Walk(test.Input, func(s string) {
				got = append(got, s)
			})

			if !slices.Equal(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
