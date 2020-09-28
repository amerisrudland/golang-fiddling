package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	response := Repeat("a", 5)
	expected := "aaaaa"

	if response != expected {
		t.Errorf("expected %q but got %q", expected, response)
	}
}
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("b", 4)
	fmt.Println(repeated)
	// Output: bbbb
}
