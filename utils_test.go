package main

import (
	"strings"
	"testing"
)

// test  for unsupported and unallowed characters
func Test_containsUnsupportedCharacters(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    bool
		wantErr string
	}{
		{
			name:    "No unsupported characters",
			input:   "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789???",
			want:    false,
			wantErr: "",
		},
		{
			name:    "One unsupported character",
			input:   "abcde\tfghij",
			want:    true,
			wantErr: "Error: input contains non-printable character: %q\n",
		},
		{
			name:    "Multiple unsupported characters",
			input:   "abcde\a\bfghi$j",
			want:    true,
			wantErr: "Error: input contains non-printable character: %q\n",
		},
		{
			name:  "Empty string",
			input: "",
			want:  false,
			wantErr: `Usage: go run . [OPTION] [STRING] [BANNER]

			Example: go run . --align=right something standard`,
		},
		{
			name:    "String with only unsupported characters",
			input:   "\a\b\t\v",
			want:    true,
			wantErr: "Error: input contains non-printable character: %q\n",
		},
		{
			name:    "String with special characters",
			input:   "Hello, world!How are you?",
			want:    false,
			wantErr: "",
		},

		{
			name:    "string with unicode characters",
			input:   "こんにちは、世界！",
			want:    true,
			wantErr: "Error: input contains unallowed character: %q\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, errmsg := containsUnsupportedCharacters(tt.input)
			if got != tt.want {
				t.Errorf("containsUnsupportedCharacters() got = %v, want %v", got, tt.want)
			}
			if got && !strings.Contains(errmsg, "Error: input contains non-printable/ unallowed character") != tt.want {
				t.Errorf("containsUnsupportedCharacters() returned unexpected error message: %s,", errmsg)
			}
		})
	}
}

// test for valid banner flag
func Test_validBanner(t *testing.T) {
	tests := []struct {
		name   string
		banner string
		want   bool
	}{
		{
			name:   "Valid banner 'standard'",
			banner: "standard",
			want:   true,
		},
		{
			name:   "Valid banner 'shadow'",
			banner: "shadow",
			want:   true,
		},
		{
			name:   "Valid banner 'thinkertoy'",
			banner: "thinkertoy",
			want:   true,
		},
		{
			name:   "Invalid banner 'other'",
			banner: "other",
			want:   false,
		},
		{
			name:   "Empty banner",
			banner: "",
			want:   false,
		},
		{
			name:   "Uppercase banner 'Standard'",
			banner: "Standard",
			want:   false,
		},
		{
			name:   "Mixed case banner 'ThinKerToy'",
			banner: "ThinKerToy",
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validBanner(tt.banner); got != tt.want {
				t.Errorf("validBanner() = %v, want %v", got, tt.want)
			}
		})
	}
}
