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
		acctual := cleanInput(c.input)
		if len(acctual) != len(c.expected) {
			t.Errorf("Length does not match '%v' vs '%v'", acctual, c.expected)
		}

		for i := range acctual {
			word := acctual[i]
			expected := c.expected[i]
			if word != expected {
				t.Errorf("failed")
			}
		}
	}

}
