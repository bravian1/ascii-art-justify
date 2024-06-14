package main

import (
	"fmt"
	"strings"
)

func printJustify(input string, asciiMap map[rune][]string) {
	input = strings.ReplaceAll(input, "\\n", "\n")
	words := strings.Split(input, "\n")
	for _, word := range words {
		justifyHelper(word, asciiMap)
	}
}

func justifyHelper(input string, asciiMap map[rune][]string) {
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
