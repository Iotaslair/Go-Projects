package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("Test Repeat with 0 items", func(t *testing.T) {
		got := Repeat("a", 0)
		want := ""

		if got != want {
			t.Errorf("Got %q Want %q", got, want)
		}
	})

	t.Run("Repeat 5 times", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"

		if got != want {
			t.Errorf("Got %q Want %q", got, want)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	//Output: aaaaa
}
