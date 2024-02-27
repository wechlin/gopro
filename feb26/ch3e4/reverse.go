package reverse

import "slices"

func Reverse(s string) string {
	content := []rune(s)

	slices.Reverse(content)

	return string(content)
}
