package main

import (
	"log"
	"os"
	"strings"

	printArt "ascii-art-justify/print"
	"ascii-art-justify/utils"
)

// fetch the command line arguments,
// validate the arguments and get filepath for the banner file,
// create map from banner file,
// align content accordingly and display results to user
func main() {
	args := os.Args[1:]
	bannerfile, flag, input := utils.ValidateArgs(args)
	if validInput, offendingCharacter := utils.IsValidInput(input); !validInput {
		log.Fatalf("Error: input contains unallowed character: %q\n", offendingCharacter)
	}

	asciiMap := utils.CreateMap(bannerfile)
	if asciiMap == nil {
		os.Exit(1)
	}
	data := strings.ReplaceAll(input, "\\n", "\n")
	words := strings.Split(data, "\n")

	for _, word := range words {
		if flag == "justify" {
			printArt.Justify(word, asciiMap)
		} else {
			printArt.Align(word, flag, asciiMap)
		}
	}
}
