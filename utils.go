package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func getTerminalWidth() int {
	cmd := exec.Command("stty", "size")
	// command.Run()
	cmd.Stdin = os.Stdin
	output, err := cmd.Output()
	// fmt.Println("Output: ", output)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	sizeString := string(output)
	sizestring := strings.Split(sizeString, " ")

	size, err := strconv.Atoi(strings.Trim(sizestring[1], "\n"))
	if err != nil {
		log.Fatal("Error: ", err)
	}
	// fmt.Println(size)
	return size
}

func getSpacesBetween(flag string, asciiString string) int {
	terminalWidth := getTerminalWidth()

	spaces := 0
	switch flag {
	case "right":
		spaces = terminalWidth - len(asciiString)
	case "left":
		spaces = 0
	case "center":
		spaces = (terminalWidth - len(asciiString)) / 2
	}
	return spaces
}

func validateArgs(args []string) (string, string, string) {
	var shouldAlign bool
	var userInput string
	var flag string
	bannerfile := "standard"

	if len(args) == 4 {
		if flag, shouldAlign = checkFlag(args[1]); shouldAlign {
			userInput = args[2]
			bannerfile = args[3]
		} else {
			printErrorAndExit()
		}
	} else if len(args) == 3 {
		if flag, shouldAlign = checkFlag(args[1]); shouldAlign {
			userInput = args[2]
		} else {
			userInput = args[1]
			if validBanner(args[2]) {
				bannerfile = args[2]
			} else {
				printErrorAndExit()
			}
		}
	} else if len(args) == 2 {
		userInput = args[1]
		if strings.HasPrefix(userInput, "--align=") {
			printErrorAndExit()
		}
	} else {
		printErrorAndExit()
	}

	if len(userInput) == 0 {
		printErrorAndExit()
	}
	return bannerfile, flag, userInput
}

// check flag
func checkFlag(input string) (string, bool) {
	if strings.HasPrefix(input, "--align=") {
		s := strings.Split(input, "=")
		// fmt.Printf("%q\n", s)
		if !((s[1] == "left") || (s[1] == "right") || (s[1] == "justify") || (s[1] == "center")) {
			printErrorAndExit()
		} else {
			flagtype := strings.Trim(input, "-align=")
			return flagtype, true
		}
	}
	return "", false
}

// Print error message
func printErrorAndExit() {
	fmt.Printf("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard\n")
	os.Exit(0)
}

// valid banner
func validBanner(banner string) bool {
	return banner == "standard" || banner == "shadow" || banner == "thinkertoy"
}

func containsUnsupportedCharacters(input string) (bool, string) {
	// special characters
	NonPrintableChars := []string{"\\a", "\\b", "\\t", "\\v", "\\f", "\\r", "\a", "\b", "\t", "\v", "\f", "\r"}
	for _, char := range NonPrintableChars {
		if contains := strings.Contains(input, char); contains {
			errmsg := fmt.Sprintf("Error: input contains non-printable character: %q\n", char)
			return true, errmsg
		}
	}
	// other characters
	input = strings.ReplaceAll(input, "\\n", "\n")
	for _, ch := range input {
		if !((ch >= 32 && ch <= 126) || ch == '\n') {
			errmsg := fmt.Sprintf("Error: input contains unallowed character: %q\n", ch)
			return true, errmsg
		}
	}
	return false, ""
}
