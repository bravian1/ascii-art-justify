package utils

import (
	"strings"
	"testing"
)

// test  for unsupported and unallowed characters
func Test_IsValidInput(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    bool
		wantErr string
	}{
		{
			name:    "No unsupported characters",
			input:   "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789???",
			want:    true,
			wantErr: "",
		},
		{
			name:    "One unsupported character",
			input:   "abcde\tfghij",
			want:    false,
			wantErr: "Error: input contains non-printable character: %q\n",
		},
		{
			name:    "Multiple unsupported characters",
			input:   "abcde\a\bfghi$j",
			want:    false,
			wantErr: "Error: input contains non-printable character: %q\n",
		},
		{
			name:  "Empty string",
			input: "",
			want:  true,
			wantErr: `Usage: go run . [OPTION] [STRING] [BANNER]

			Example: go run . --align=right something standard`,
		},
		{
			name:    "String with only unsupported characters",
			input:   "\a\b\t\v",
			want:    false,
			wantErr: "Error: input contains non-printable character: %q\n",
		},
		{
			name:    "String with special characters",
			input:   "Hello, world!How are you?",
			want:    true,
			wantErr: "",
		},

		{
			name:    "string with unicode characters",
			input:   "こんにちは、世界！",
			want:    false,
			wantErr: "Error: input contains unallowed character: %q\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, errmsg := IsValidInput(tt.input)
			if got != tt.want {
				t.Errorf("ContainsUnsupportedCharacters() got = %v, want %v", got, tt.want)
			}
			if got && !strings.Contains(errmsg, "Error: input contains non-printable/ unallowed character") != tt.want {
				t.Errorf("ContainsUnsupportedCharacters() returned unexpected error message: %s,", errmsg)
			}
		})
	}
}

// test for valid banner
func Test_ValidBanner(t *testing.T) {
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
			if got := ValidBanner(tt.banner); got != tt.want {
				t.Errorf("ValidBanner() = %v, want %v", got, tt.want)
			}
		})
	}
}

// test for valid arguments depending on the length of arguments
func Test_ValidateArgs(t *testing.T) {
	tests := []struct {
		name          string
		args          []string
		expectedFile  string
		expectedFlag  string
		expectedInput string
	}{
		{
			name:          "Valid case with all arguments",
			args:          []string{"--align=justify", "Hello world", "shadow"},
			expectedFile:  "shadow",
			expectedFlag:  "justify",
			expectedInput: "Hello world",
		},
		{
			name:          "Valid case with flag and userInput",
			args:          []string{"--align=right", "Hello World"},
			expectedFile:  "standard",
			expectedFlag:  "right",
			expectedInput: "Hello World",
		},
		{
			name:          "Valid case with userInput and bannerfile",
			args:          []string{"Hello World", "shadow"},
			expectedFile:  "shadow",
			expectedFlag:  "",
			expectedInput: "Hello World",
		},
		{
			name:          "Valid case with only userInput",
			args:          []string{"user_input"},
			expectedFile:  "standard",
			expectedFlag:  "",
			expectedInput: "user_input",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := ValidateArgs(tt.args)
			if got != tt.expectedFile {
				t.Errorf("ValidateArgs() got = %v, want %v", got, tt.expectedFile)
			}
			if got1 != tt.expectedFlag {
				t.Errorf("ValidateArgs() got1 = %v, want %v", got1, tt.expectedFlag)
			}
			if got2 != tt.expectedInput {
				t.Errorf("ValidateArgs() got2 = %v, want %v", got2, tt.expectedInput)
			}
		})
	}
}

// test to check flags for alignment flag
func Test_CheckFlag(t *testing.T) {
	tests := []struct {
		name       string
		flagString string
		want       string
		want1      bool
		wantErr    bool
	}{
		{
			name:       "Valid left alignment flag",
			flagString: "--align=left",
			want:       "left",
			want1:      true,
			wantErr:    false,
		},
		{
			name:       "Valid right alignment flag",
			flagString: "--align=right",
			want:       "right",
			want1:      true,
			wantErr:    false,
		},
		{
			name:       "Valid justify alignment flag",
			flagString: "--align=justify",
			want:       "justify",
			want1:      true,
			wantErr:    false,
		},
		{
			name:       "Valid center alignment flag",
			flagString: "--align=center",
			want:       "center",
			want1:      true,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := CheckFlag(tt.flagString)
			if got != tt.want {
				t.Errorf("CheckFlag() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CheckFlag() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}