package xstring

import "strings"

// BuilderConcat - concat str by go pkg string.Builder
func BuilderConcat(list ...string) string {
	b := strings.Builder{}
	l := 0
	for _, v := range list {
		l += len(v)
	}
	b.Grow(l)
	for _, v := range list {
		b.WriteString(v)
	}
	return b.String()
}
