package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "UPPERWORD lowerword",
			expected: []string{"upperword", "lowerword"},
		},
		{
			input:    "h eLl O WorLD",
			expected: []string{"h", "ell", "o", "world"},
		},
		{
			input:    "",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Error length! Expected:%d || Actual: %d", len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Incorect word! Expected:%s || Actual: %s", expectedWord, word)
			}
		}
	}
}
