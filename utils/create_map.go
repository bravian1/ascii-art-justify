package utils

import (
	"bufio"
	"fmt"
	"os"
)

// program to create a map of the 95 printable ascii characters to their art representation from a banner file
func CreateMap(file *os.File) map[rune][]string {
	// use scanner to avoid issues with thinkertoy.txt
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	asciiMap := make(map[rune][]string)
	letter := ' '
	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		if count != 8 {
			asciiMap[letter] = append(asciiMap[letter], line)
			count++

		} else {
			count = 0
			letter++
		}
	}
	if (len(asciiMap['A'])) != 8 {
		fmt.Println("the banner file does not contain the expected format\n", file)
		return nil
	}
	if len(asciiMap) == 0 || len(asciiMap) != 95 {
		fmt.Println("the banner file does not contain the expected format\n", file)
		return nil
	}
	return asciiMap
}
