package helpers

import (
	. "github.com/onsi/gomega"
)

// ExpectMissingSubstrings checks for existence of one or many substrings
func ExpectSubstrings(data any, substrings ...string) {
	for _, s := range substrings {
		Expect(data).To(ContainSubstring(s))
	}
}

// ExpectMissingSubstrings checks for non-existence of one or many substrings
func ExpectMissingSubstrings(data any, substrings ...string) {
	for _, s := range substrings {
		Expect(data).ToNot(ContainSubstring(s))
	}
}
