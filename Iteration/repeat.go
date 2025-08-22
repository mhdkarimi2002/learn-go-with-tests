package iteration

import "strings"

func Repeat(character string, repeatCount int) string {
	var builder strings.Builder
	builder.Grow(len(character) * repeatCount) // preallocate for performance
	for i := 0; i < repeatCount; i++ {
		builder.WriteString(character)
	}
	return builder.String()
}
