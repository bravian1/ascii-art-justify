package asciiart

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// isValid checks if the input string contains only acceptable characters. ([32,126] & [8,10])
// The special characters newline (10), carriage return (13), and backspace (8).
// - bool: true if contains only acceptable, false otherwise.
func isValidInput(args string) (bool, rune) {
	for _, ch := range args {
		if !((ch >= 32 && ch <= 126) || (ch >= 8 && ch <= 10)) {
			return false, ch
		}
	}
	return true, ' '
}

/*
* Replace all special characters with characters that can be recognized and processed by golang
* Make them fit inside a rune
 */
func replaceSpecialcharacters(s string) string {
	replacer := strings.NewReplacer(
		"\\v", "\\n\\n\\n\\n",
		"\n", "\\n",
		"\\t", "    ",
		"\\b", "\b",
		"\\r", "\r",
		"\\a", "\a",
		"\\f", "\f",
	)
	return replacer.Replace(s)
}

// printAscii prints the ascii art representation of the given word using the given alphabet
func printAscii(word string, alphabet map[rune][]string) {
	for i := 0; i < CHARACTER_HEIGHT; i++ {
		lineOutput := ""
		for _, l := range word {
			switch l {
			case '\n':
				fmt.Println()
			default:
				lineOutput += alphabet[rune(l)][i]
			}
		}

		fmt.Println(lineOutput)
	}
}

// printAscii prints the ascii art representation of the given word using the given alphabet
func printAsciiAlign(word string, alphabet map[rune][]string, flagtype string) {
	for i := 0; i < CHARACTER_HEIGHT; i++ {
		lineOutput := ""
		for _, l := range word {
			switch l {
			case '\n':
				fmt.Println()
			default:
				lineOutput += alphabet[rune(l)][i]
			}
		}
		padding, _ := getsize(lineOutput, flagtype)

		f := strings.Repeat(" ", padding) + lineOutput
		// fmt.Printf(fmt.Sprintf("%%-%ds", width/2), fmt.Sprintf(fmt.Sprintf("%%%ds", width), lineOutput))
		fmt.Println(f)
	}
}

/*
we get terminal width
split the word according to spaces so as to get the number of spaces to be added between the words
get the total length of the ascii characters of the word
total spaces to be put is terminal width - total length of the ascii characters of the word
if there are no spaces or the terminal width is less than the total length of the ascii characters of the word then it should print normally
else padding (spacebtwnwords) is calculated as total spaces/number of spaces
we -1 to accomodate cat -e if required
extraspace is calculated as totalspaces%number of spaces
we print the ascii characters of the word
if we are not at the end of the word then we  add the padding and distribute the extra spaces evenly
*/
func printAsciiJustify(word string, alphabet map[rune][]string, flagtype string) {
	// get terminal width
	_, width := getsize(word, flagtype)

	words := strings.Split(word, " ")
	totallength := 0
	for _, word := range words {
		for _, l := range word {
			totallength += len(alphabet[l][0])
		}
	}

	gaps := len(words) - 1
	totalspaces := width - totallength
	if gaps == 0 || totalspaces <= 0 {
		printAscii(word, alphabet)
		return
	}
	spacebtwnwords := (totalspaces / gaps)
	extraspaces := totalspaces % gaps

	for i := 0; i < CHARACTER_HEIGHT; i++ {
		lineOutput := ""
		for j, word := range words {
			for _, l := range word {
				lineOutput += alphabet[l][i]
			}
			if j < gaps {

				lineOutput += strings.Repeat(" ", spacebtwnwords)
				if j < extraspaces {
					lineOutput += " "
				}
			}
		}
		fmt.Println(lineOutput)
	}
}

func getsize(v string, flagtype string) (int, int) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	s := string(output)
	ss := strings.Split(s, " ")
	width, err := strconv.Atoi(strings.Trim(ss[1], "\n"))
	if err != nil {
		log.Fatal(err)
	}
	padding := 0
	switch flagtype {

	case "left":
		padding = 0
	case "right":
		padding = (width - len(v))
	case "center":
		padding = (width - len(v)) / 2

	default:
		padding = 0
	}

	return padding, width
}

/*
* Processes the input array and extracts whatever the program needs
* It is also responsible of dealing with optional parameters like the file names
* Exits if input does not meet the expected format
 */
func processInput(arr []string) (bool, string) {
	if len(arr) < 2 || len(arr) >= 5 {
		return false, ""
	}

	filename := "standard.txt"
	if len(arr) == 4 {
		switch arr[3] {
		case "shadow", "shadow.txt":
			filename = "shadow.txt"
		case "thinkertoy", "thinkertoy.txt":
			filename = "thinkertoy.txt"
		}
	}
	if len(arr) == 3 {
		switch arr[2] {
		case "shadow", "shadow.txt":
			filename = "shadow.txt"
		case "thinkertoy", "thinkertoy.txt":
			filename = "thinkertoy.txt"
		}
	}
	return true, filename
}

func processBackspace(s string) string {
	if len(s) > 1 {
		temp := ""
		for index, ch := range s {
			if ch != '\b' {
				temp += string(ch)
			} else {
				temp = temp[:index-1]
			}
		}
		return temp
	} else {
		return ""
	}
}
