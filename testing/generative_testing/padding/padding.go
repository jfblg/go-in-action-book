package padding

import (
	"log"
	"strings"
)

func Pad(s string, max uint) string {
	log.Printf("Testing len: %d, Str: %s", max, s)
	len := uint(len(s))
	if len > max {
		return s[:max]
	}
	s += strings.Repeat(" ", int(max-len))

	return s
}
