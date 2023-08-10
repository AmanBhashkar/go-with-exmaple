package integers

import (
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("Adding two +ve integer", func(t *testing.T) {
		got := Add(5, 10)
		want := 15
		assertAdd(t, got, want)
	})
	t.Run("Adding two -ve integer", func(t *testing.T) {
		got := Add(-2, -4)
		want := -6
		assertAdd(t, got, want)
	})
}

func assertAdd(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
