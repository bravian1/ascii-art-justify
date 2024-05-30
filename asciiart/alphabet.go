package asciiart

import (
	"bufio"
	"fmt"
	"os"
)

/*
* CreateAlphabet(string) map[rune][]string
*
* reads a file and creates a map from letters to their representations.
* input:
* 	filename // the name of the file holding the letter representation
* output:
*	a map from letter in rune form to the letter representation as []string
 */
func CreateAlphabet(filename string) (bool, map[rune][]string) {
	file, err := os.Open(filename)
	if err != nil {
		return false, nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// skip first line
	scanner.Scan()
	alphabet := make(map[rune][]string)
	var letter rune = ' '
	var count int
	lines := 0
	for scanner.Scan() {
		lines++
		if count != CHARACTER_HEIGHT {
			alphabet[letter] = append(alphabet[letter], scanner.Text())
			count++
		} else {
			letter++
			count = 0
		}
		if err := scanner.Err(); err != nil {
			fmt.Printf("[Internal Error] Error reading file: %v", err)
			os.Exit(0)
		}
	}
	// does the alphabet created from the file contain all the allowed characters
	if len(alphabet) != NUMOFASCII_XTERS {
		fmt.Printf("[Error] Expected %d but got %d characters in the alphabet. Ensure you have the correct ascii art file\n", NUMOFASCII_XTERS, len(alphabet))
		os.Exit(0)
	} else if lines != NUMOFLINESINBANNER {
		fmt.Printf("[Error] Expected %d but got %d lines in the banner file. Ensure you have the correct ascii art file length\n", NUMOFLINESINBANNER, lines)
		os.Exit(0)
	}
	return true, alphabet
}
