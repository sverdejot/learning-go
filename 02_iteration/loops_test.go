package loops

import "testing"
import "fmt"

func TestRepeat(t *testing.T) {
	t.Run("repeat with given values", func(t *testing.T) {
		repeated := Repeated("a", 5)
		expected := "aaaaa"

		assertValuesAreEqual(t, expected, repeated)
	})

	t.Run("repeat with no cycles", func(t *testing.T) {
		repeated := Repeated("a", 0)
		expected := ""

		assertValuesAreEqual(t, expected, repeated)
	})

}

func assertValuesAreEqual(t testing.TB, expected, got string) {
	t.Helper()
	if got != expected {
		t.Errorf("expected %q but got %q", expected, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeated("a", 5)
	}
}

func ExampleRepeated() {
	result := Repeated("a", 5)
	fmt.Println(result)
	// Output: aaaaa
}
