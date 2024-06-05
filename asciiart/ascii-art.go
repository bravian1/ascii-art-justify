package asciiart

import (
	"fmt"
	"os"
	"strings"
)

//echo -e "lines\ncols" | tput -S for getting terminal width and height

const (
	CHARACTER_HEIGHT   = 8
	NUMOFASCII_XTERS   = 95
	NUMOFLINESINBANNER = 854
)

// Run processes an array of strings to generate ASCII art from a specified file.
func Run(arr []string) {
	if len(arr) < 2 {
		printUsageMessageAndExit()
		return
	}

	
	inputstring, flagtype:=checkflag(arr)
	
	processed, filename := processInput(arr)
	if !processed {
		printUsageMessageAndExit()
		return
	}

	args := replaceSpecialcharacters(inputstring)
	args = processBackspace(args)

	if valid, xter := isValidInput(args); !valid {
		fmt.Printf("[Error] Input contains unacceptable character %q\n", xter)
		return
	}

	if args == "\\n" {
		fmt.Println()
		return
	} else if args == "" {
		return
	} else {
		input := strings.Split(args, "\\n")
		created, alphabet := CreateAlphabet(filename)
		if !created {
			fmt.Printf("Could not create the alphabet. Are you sure %s exists and is a valid ascii file?\n", filename)
			os.Exit(0)
		}
		

		for _, word := range input {
			if word == "" {
				fmt.Println()
			} else if flagtype == "justify" {
				printAsciiJustify(word, alphabet, flagtype)
			} else if flagtype == "right" || flagtype == "center" || flagtype == "left" {
				printAsciiAlign(word, alphabet, flagtype)
			} else {
				printAscii(word, alphabet)
			}
		}
	}
}

// printUsageMessageAndExit prints the usage message and exits
func printUsageMessageAndExit() {
	fmt.Println(`
		Usage: go run . [OPTION] [STRING]
		EX: go run . --align=<align> "something"
	`)
	os.Exit(0)
}

func checkflag(arr []string) (string, string) {
	var inputstring, flagtype string
	
	if strings.HasPrefix(arr[1], "--align=") {
		flagtype = strings.TrimPrefix(arr[1], "--align=")
		if len(arr) == 3 {
			inputstring = arr[2]
		} else if len(arr) == 4 {
			inputstring=arr[2]
			
		}
	} else {
		inputstring = arr[1]
		return inputstring, ""
	}

	return inputstring, flagtype
}
