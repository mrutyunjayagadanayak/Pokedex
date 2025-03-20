package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " Hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " Hello World ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " HelloWorld ",
			expected: []string{"helloworld"},
		},
		{
			input:    "",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Expected %d words, got %d words", len(c.expected), len(actual))
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				fmt.Print(word)
				t.Errorf("Expected - %s. Got - %s", expectedWord, word)
			}
		}
	}
}
