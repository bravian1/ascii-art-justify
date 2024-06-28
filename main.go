package main

import (
	"fmt"
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

	file, err := os.Open(bannerfile + ".txt")
	if err != nil {
		utils.GetFile(bannerfile + ".txt")
		file, err = os.Open(bannerfile + ".txt")
		if err != nil {
			fmt.Println("error: ", err)
		}
	}
	defer file.Close()

	asciiMap := utils.CreateMap(file)
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