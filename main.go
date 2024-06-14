package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	// get banner file
	// turn banner file to map
	// start by opening file

	bannerfile, flag, input := validateArgs(os.Args)
	// fmt.Println(bannerfile)
	// fmt.Println(flag)
	// fmt.Println(input)
	file, err := os.Open(bannerfile + ".txt")
	if err != nil {
		GetFile(bannerfile + ".txt")
		file, _ = os.Open(bannerfile + ".txt")
		// log.Fatal(err)
	}
	defer file.Close()

	asciiMap := createMap(file)
	if flag == "justify" {
		printJustify(input, asciiMap)
		return
	}
	printAlign(input, flag, asciiMap)
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

func createMap(file *os.File) map[rune][]string {
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	asciiMap := make(map[rune][]string)
	letter := ' '
	count := 0
	line := ""
	for scanner.Scan() {
		line = scanner.Text()
		if count != 8 {
			asciiMap[letter] = append(asciiMap[letter], line)
			count++

		} else {
			count = 0
			letter++
		}
	}
	fmt.Println(len(line))
	fmt.Println(len(asciiMap))
	if len(asciiMap) == 0 || len(asciiMap) != 95 {
		fmt.Println("the banner file   does not contain the expected format\n", file)
		return nil
	}

	return asciiMap
}

func printNormal(input string, asciiMap map[rune][]string) {
	input = strings.ReplaceAll(input, "\\n", "\n")
	inputSlice := strings.Split(input, "\n")
	for _, line := range inputSlice {
		if line == "" {
			fmt.Println()
		} else {
			for i := 0; i < 8; i++ {
				lineOutput := ""

				for _, char := range line {
					lineOutput += asciiMap[char][i]
				}

				fmt.Println(lineOutput)
			}
		}
		fmt.Println()
	}
}

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

func printAlign(input string, flag string, asciiMap map[rune][]string) {
	input = strings.ReplaceAll(input, "\\n", "\n")
	inputSlice := strings.Split(input, "\n")
	combinedWidth := 0
	for _, word := range inputSlice {
		for _, char := range word {
			combinedWidth += len(asciiMap[char][0])
		}
	}
	// get terminal width
	terminalWidth := getTerminalWidth()
	if combinedWidth > terminalWidth {
		printNormal(input, asciiMap)
		return
	}

	for _, word := range inputSlice {
		if word == "" {
			fmt.Println()
		} else {
			for i := 0; i < 8; i++ {
				lineOutput := ""
				for _, char := range word {

					line, ok := asciiMap[char]
					if !ok {
						fmt.Printf("Unavailable %c\n", char)
					}
					lineOutput += line[i]
				}
				spaces := getSpacesBetween(flag, lineOutput)
				lineOutput = strings.Repeat(" ", spaces) + lineOutput
				fmt.Println(lineOutput)
			}
		}
	}
}

func printJustify(input string, asciiMap map[rune][]string) {
	// split the word
	splitWord := strings.Split(input, " ")
	// get the number of gaps
	numGaps := len(splitWord) - 1
	// get the combined width of the words as ascii
	combinedWidth := 0
	for _, word := range splitWord {
		for _, char := range word {
			combinedWidth += len(asciiMap[char][0])
		}
	}
	// get terminal width
	terminalWidth := getTerminalWidth()
	if combinedWidth > terminalWidth || numGaps == 0 {
		printNormal(input, asciiMap)
		return
	}
	// calculate the total number of spaces
	totalSpaces := terminalWidth - combinedWidth
	// divide the spaces equally depending on the number of words
	spacesBetween := totalSpaces / numGaps
	// print
	for i := 0; i < 8; i++ {
		lineOutput := ""
		for j, word := range splitWord {
			for _, char := range word {
				lineOutput += asciiMap[char][i]
			}
			if j < numGaps {
				lineOutput += strings.Repeat(" ", spacesBetween)
			}
		}

		fmt.Println(lineOutput)
	}
}
