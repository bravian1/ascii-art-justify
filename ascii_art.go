package main

import (
	"fmt"
	"strings"
)

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
