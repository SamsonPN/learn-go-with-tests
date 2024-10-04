package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("repeat 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		assertCorrectRepeat(t, repeated, expected)
	})
	t.Run("repeat n times", func(t *testing.T) {
		repeated := Repeat("a", 3)
		expected := "aaa"

		assertCorrectRepeat(t, repeated, expected)
	})
}

func assertCorrectRepeat(t testing.TB, repeated, expected string) {
	t.Helper()
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

// Benchmarks
// very similar to writing tests, i.e. always start with Benchmark[func-name]
// testing.B gives you access to b.N which will run a function N times
// by default, Benchmarks are run SEQUENTIALLY!

// to run this, do go test-bench=. in the directory where the benchmark should be
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
