package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "charmander bulbasaur pikachu", 
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input: "", 
			expected: []string{},
		},
		{
			input: "GARCHOMP eevEE Vaporeon", 
			expected: []string{"garchomp", "eevee", "vaporeon"},
		},
		{
			input: "  GARCHOMP eevEE Vaporeon  ",
			expected: []string{"garchomp", "eevee", "vaporeon"},
		},
		{
			input: "  GARCHOMP  eevEE   Vaporeon  ", 
			expected: []string{"garchomp", "eevee", "vaporeon"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		
		// Check slice lengths
		if len(actual) != len(c.expected) {
			t.Errorf("Slice lengths do not match. \n\nExpected Slice: %v\nActual Slice: %v\n\nExpected Length: %v\nActual Length: %v\n\n---------------------------------", c.expected, actual, len(c.expected), len(actual))
			continue
		}
		
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			
			// Check each word with the expected output
			if word != expectedWord {
				t.Errorf("Output mismatch. \n\nExpected Slice: %v\nActual Slice: %v\n\nExpected Word: %v\nActual Word: %v\n\n---------------------------------", c.expected, actual, word, expectedWord)
			}
		}
	}
}