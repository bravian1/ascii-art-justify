package printArt

import (
	"fmt"
)

func Normal(input string, asciiMap map[rune][]string) {
	if input == "" {
		fmt.Println()
	} else {
		for i := 0; i < 8; i++ {
			lineOutput := ""
			for _, char := range input {
				lineOutput += asciiMap[char][i]
			}
			fmt.Println(lineOutput)
		}
	}
}