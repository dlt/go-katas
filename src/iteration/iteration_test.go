package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestRevert(t *testing.T) {
	text := Revert("aaab")
	expected := "baaa"

	if text != expected {
		t.Errorf("got %q expected %q", text, expected)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("foo")
	}
}