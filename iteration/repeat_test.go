package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	testWithParams := func(t *testing.T, char string, repeatNum int) (string, string) {
		t.Helper()
		response := Repeat(char, repeatNum)
		expected := strings.Repeat(char, repeatNum)
		return response, expected
	}

	t.Run("test with EqualFold", func(t *testing.T) {
		response, expected := testWithParams(t, "a", 4)
		if !strings.EqualFold(response, expected) {
			t.Errorf("expected %q but got %q", expected, response)
		}
	})

	t.Run("test with Compare", func(t *testing.T) {
		response, expected := testWithParams(t, "b", 5)
		if strings.Compare(response, expected) != 0 {
			t.Errorf("expected %q but got %q", expected, response)
		}
	})

	t.Run("test with Count", func(t *testing.T) {
		char := "c"
		response, expected := testWithParams(t, char, 6)
		if strings.Count(response, char) != strings.Count(expected, char) {
			t.Errorf("expected %d chars but got %d chars", strings.Count(expected, char), strings.Count(response, char))
		}
	})
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
