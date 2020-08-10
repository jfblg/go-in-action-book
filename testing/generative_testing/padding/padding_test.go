package padding

import (
	"testing"
	"testing/quick"
)

func TestPad(t *testing.T) {
	if r := Pad("test", 6); len(r) != 6 {
		t.Errorf("Expected len 6, got %d", len(r))
	}
}

func TestPadGenerative(t *testing.T) {
	fn := func(s string, max uint8) bool {
		p := Pad(s, uint(max))
		return len(p) == int(max)
	}
	if err := quick.Check(fn, &quick.Config{MaxCount: 500}); err != nil {
		t.Error(err)
	}
}
