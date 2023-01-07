package dbfs

import (
	"os"
	"strings"
)

// Path concatenates part of path using os.PathSeparator
func Path(parts ...string) string {
	builder := strings.Builder{}
	for _, part := range parts[:len(parts)-1] {
		builder.WriteString(part)
		builder.WriteRune(os.PathSeparator)
	}
	builder.WriteString(parts[len(parts)-1])
	return builder.String()
}
