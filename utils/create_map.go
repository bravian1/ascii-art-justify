package utils

import (
	"bufio"
	"fmt"
	"os"
)

// program to create a map of the 95 printable ascii characters to their art representation from a banner file
func CreateMap(filename string) map[rune][]string {
	file, err := os.Open(filename + ".txt")
	if err != nil {
		// GetFile(bannerfile + ".txt")
		// file, err = os.Open(filename + ".txt")
		// if err != nil {
		// 	fmt.Println("error: ", err)
		// }
		fmt.Printf("Error opening file %q\n", filename)
		return nil
	}
	defer file.Close()

	// use scanner to avoid issues with thinkertoy.txt
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	asciiMap := make(map[rune][]string)
	letter := ' '
	count := 0
	numLinesInFile := 0

	for scanner.Scan() {
		line := scanner.Text()
		numLinesInFile++
		if count != 8 {
			asciiMap[letter] = append(asciiMap[letter], line)
			count++

		} else {
			count = 0
			letter++
		}
	}
	if (len(asciiMap['A'])) != 8 {
		fmt.Printf("the banner file %q does not contain the expected format\n", filename + ".txt")
		return nil
	}
	if len(asciiMap) == 0 || len(asciiMap) != 95 || numLinesInFile != 854 {
		fmt.Printf("the banner file %q does not contain the expected format\n", filename + ".txt")
		return nil
	}
	return asciiMap
}
