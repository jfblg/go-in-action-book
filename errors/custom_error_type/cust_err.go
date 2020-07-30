package parseerror

import "fmt"

// ParseError is a custom error type
type ParseError struct {
	Message    string
	Line, Char int
}

func (p *ParseError) Error() string {
	format := "%s on Line %d, Char %d"
	return fmt.Sprintf(format, p.Message, p.Line, p.Char)
}
