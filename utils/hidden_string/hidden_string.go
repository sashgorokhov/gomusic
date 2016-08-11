package hidden_string

import "strings"

type HiddenString string

func (s HiddenString) String() string {
	return HideString(s.Original())
}

func (s HiddenString) Original() string {
	return string(s)
}

func HideString(s string) string {
	return strings.Repeat("*", len(s))
}
