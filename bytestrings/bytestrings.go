package bytestrings

import (
	"bytes"
	"unicode"
)

func ToUpper(b byte) byte {
	return byte(unicode.ToUpper(rune(b)))
}

func IsSpace(b byte) bool {
	return unicode.IsSpace(rune(b))
}

func CompareCaseInsensitive(a []byte, b []byte) int {
	a = bytes.ToLower(a)
	b = bytes.ToLower(b)
	return bytes.Compare(a, b)
}
