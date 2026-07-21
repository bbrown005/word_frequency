package main

import (
	"reflect"
	"testing"
)

func TestCountWords(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]int
	}{
		{
			name:  "repeated words",
			input: "hello world hello",
			expected: map[string]int{
				"hello": 2,
				"world": 1,
			},
		},
		{
			name:     "empty string",
			input:    "",
			expected: map[string]int{},
		},
		{
			name:  "single word",
			input: "hello",
			expected: map[string]int{
				"hello": 1,
			},
		},
		{
			name:  "multiple spaces",
			input: "hello    world  hello",
			expected: map[string]int{
				"hello": 2,
				"world": 1,
			},
		},
		{
			name:  "capitalized words",
			input: "HELLO world Hello",
			expected: map[string]int{
				"Hello": 1,
				"world": 1,
				"HELLO": 1,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := CountWords(test.input)

			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("got %v, want %v", result, test.expected)
			}
		})
	}
}
