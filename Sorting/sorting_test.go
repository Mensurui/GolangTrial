package main

import (
	"sort"
	"testing"
)

func TestStringSlice(t *testing.T) {
	var tests = []struct {
		input []string
		want  []string
	}{
		{[]string{"Mensur", "Shawshanker"}, []string{"Mensur", "Shawshanker"}},                 // already sorted
		{[]string{"Shawshanker", "Mensur"}, []string{"Mensur", "Shawshanker"}},                 // needs sorting
		{[]string{"apple", "banana", "cherry"}, []string{"apple", "banana", "cherry"}},         // already sorted
		{[]string{"cherry", "banana", "apple"}, []string{"apple", "banana", "cherry"}},         // reverse order
		{[]string{"dog", "cat", "elephant", "bat"}, []string{"bat", "cat", "dog", "elephant"}}, // random order
	}

	for _, test := range tests {
		// Sort the input slice using StringSlice
		sort.Sort(StringSlice(test.input))

		// Compare the sorted slice with the expected output
		for i, got := range test.input {
			if got != test.want[i] {
				t.Errorf("got %q, want %q", test.input, test.want)
				break
			}
		}
	}
}
