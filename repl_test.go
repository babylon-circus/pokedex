package main

import (
	"testing"
)

func TestCleanupInput(t *testing.T) {

	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "Hello World",
			expected: []string{"hello", "world"},
		},
		{
			input:    "   Hello World",
			expected: []string{"hello", "world"},
		},
		{
			input:    "   Hello World!",
			expected: []string{"hello", "world!"},
		},
		{
			input:    "   Hello WOrld!   %",
			expected: []string{"hello", "world!", "%"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Length does not match '%v' vs '%v'", actual, c.expected)
		}

		for i := range actual {
			word := actual[i]
			expected := c.expected[i]
			if word != expected {
				t.Errorf("failed")
			}
		}
	}

}
