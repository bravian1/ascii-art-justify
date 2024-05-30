package asciiart

import (
	"testing"
)

func TestIsValidInput(t *testing.T) {
	type Characters struct {
		char   string
		expect bool
	}

	table := []Characters{
		{"A", true},
		{"|", true},
		{"\r", false},
		{"\a", false},
		{"8", true},
		{"\n", true},
	}

	for _, ch := range table {
		if valid, _ := isValidInput(ch.char); valid != ch.expect {
			t.Errorf("\n[Character %q] expected %v, got %v\n", ch.char, ch.expect, valid)
		}
	}
}

func TestReplaceSpecialCharacters(t *testing.T) {
	tests := []struct {
		name             string
		input            []string
		expectedArgs     string
		expectedFilename string
	}{
		{
			name:         "Standard filename with newline",
			input:        []string{"go run .", "Hello\\nWorld"},
			expectedArgs: "Hello\\nWorld",
		},
		{
			name:         "Shadow font",
			input:        []string{"go run .", "Hello World", "shadow"},
			expectedArgs: "Hello World",
		},
		{
			name:         "Special characters replacement",
			input:        []string{"go run .", "Hello\\tWorld\\a"},
			expectedArgs: "Hello    World\a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := replaceSpecialcharacters(tt.input[1])
			if args != tt.expectedArgs {
				t.Errorf("replaceSpecialcharacters() args = %v, want %v", args, tt.expectedArgs)
			}
		})
	}
}
