package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// function to get the terminal width instead of using one fixed width
func GetTerminalWidth() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	output, err := cmd.Output()
	if err != nil {
		log.Fatal("Error: ", err)
	}
	sizeString := string(output)
	sizestring := strings.Split(sizeString, " ")

	size, err := strconv.Atoi(strings.Trim(sizestring[1], "\n"))
	if err != nil {
		log.Fatal("Error: ", err)
	}
	return size
}

// function to get spaces to add for alignment depending on the alignment flag
func GetSpaces(flag string, asciiString string) int {
	terminalWidth := GetTerminalWidth()

	spaces := 0
	switch flag {
	case "right":
		spaces = terminalWidth - len(asciiString) -1
	case "left":
		spaces = 0
	case "center":
		spaces = (terminalWidth - len(asciiString)) / 2
	}
	return spaces
}

// function to assign arguments appropriately depending on length of arguments
func ValidateArgs(args []string) (string, string, string) {
	var shouldAlign bool
	var userInput string
	var flag string
	bannerfile := "standard"

	// usage: go run . --align=right something standard
	if len(args) == 3 {
		if flag, shouldAlign = CheckFlag(args[0]); shouldAlign {
			userInput = args[1]
			bannerfile = args[2]
		} else {
			PrintErrorAndExit()
		}

		// usage: go run . --align=right something
	} else if len(args) == 2 {
		if flag, shouldAlign = CheckFlag(args[0]); shouldAlign {
			userInput = args[1]

			// usage: go run . something standard
		} else {
			userInput = args[0]
			if ValidBanner(args[1]) {
				bannerfile = args[1]
			} else {
				PrintErrorAndExit()
			}
		}

		// usage: go run . something
	} else if len(args) == 1 {
		userInput = args[0]
		if strings.HasPrefix(userInput, "--align=") {
			PrintErrorAndExit()
		}
	} else {
		PrintErrorAndExit()
	}

	if len(userInput) == 0 {
		PrintErrorAndExit()
	}
	return bannerfile, flag, userInput
}

// function to check if correct flag is passed
func CheckFlag(input string) (string, bool) {
	if strings.HasPrefix(input, "--align=") {
		flagtype := strings.TrimPrefix(input, "--align=")
		if !(flagtype == "left" || flagtype == "right" || flagtype == "center" || flagtype == "justify") {
			PrintErrorAndExit()
		} else {
			return flagtype, true
		}
	}
	return "", false
}

// function to print and exit program due to usage error
func PrintErrorAndExit() {
	fmt.Printf("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard\n")
	os.Exit(0)
}

// function to check if the correct banner is passed
func ValidBanner(banner string) bool {
	return banner == "standard" || banner == "shadow" || banner == "thinkertoy"
}

// function to check if input string contains unprintable and unsupported characters that are not within the ascii printable range
func IsValidInput(input string) (bool, string) {
	NonPrintableChars := []string{"\\a", "\\b", "\\t", "\\v", "\\f", "\\r", "\a", "\b", "\t", "\v", "\f", "\r"}
	for _, char := range NonPrintableChars {
		if contains := strings.Contains(input, char); contains {
			return false, string(char)
		}
	}
	// other characters
	for _, ch := range input {
		if !((ch >= 32 && ch <= 126) || ch == '\n') {
			return false, string(ch)
		}
	}
	return true, input
}