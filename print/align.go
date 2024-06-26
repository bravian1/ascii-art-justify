package printArt

import (
	"fmt"
	"strings"

	"ascii-art-justify/utils"
)

func Align(input string, flag string, asciiMap map[rune][]string) {
	inputSlice := strings.Split(input, "\n")
	combinedWidth := 0
	for _, word := range inputSlice {
		for _, char := range word {
			combinedWidth += len(asciiMap[char][0])
		}
	}
	// get terminal width
	terminalWidth := utils.GetTerminalWidth()
	if combinedWidth > terminalWidth {
		Normal(input, asciiMap)
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
				spaces := utils.GetSpaces(flag, lineOutput)
				lineOutput = strings.Repeat(" ", spaces) + lineOutput
				fmt.Println(lineOutput)
			}
		}
	}
}