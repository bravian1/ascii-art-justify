package main

import (
	"bufio"
	"fmt"
	"os"
)

func createMap(file *os.File) map[rune][]string {
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
	if len(asciiMap) == 0 || len(asciiMap) != 95 {
		fmt.Println("the banner file does not contain the expected format\n", file)
	}
	return asciiMap
}
