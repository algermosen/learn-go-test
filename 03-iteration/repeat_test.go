package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat character 'default'", func(t *testing.T) {
		repeated := Repeat("a", -1)
		expected := "aaaaa"

		assertRepetition(t, expected, repeated)
	})

	t.Run("repeat character 0 times", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := ""

		assertRepetition(t, expected, repeated)
	})

	t.Run("repeat character 'n' times", func(t *testing.T) {
		repeated := Repeat("a", 3)
		expected := "aaa"

		assertRepetition(t, expected, repeated)
	})
}

func assertRepetition(t testing.TB, expected, repeated string) {
	t.Helper()
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", -1)
	}
}

func ExampleRepeat() {
	// default repetiton: 5
	defaultRepeated := Repeat("a", -1)
	zeroRepeated := Repeat("a", 0)
	specificRepeated := Repeat("a", 3)
	fmt.Printf("%q %q %q", defaultRepeated, zeroRepeated, specificRepeated)
	// Output: "aaaaa" "" "aaa"
}
