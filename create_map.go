package main

import (
	"bufio"
	"os"
)

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
	return asciiMap
}
