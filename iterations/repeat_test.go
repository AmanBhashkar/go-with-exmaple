package iterations

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("na", 5)
	expected := strings.Repeat("na", 5)
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	//Output: aaaaa

}
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 7)
	}
}
