package main

import (
	"fmt"
	"log"
	"os"

	printArt "ascii-art-justify/print"
	"ascii-art-justify/utils"
)

func main() {
	// fetch the command line arguments,
	// validate the arguments and get filepath for the banner file,
	// create map from banner file,
	// align content accordingly and display results to user
	args := os.Args[1:]
	bannerfile, flag, input := utils.ValidateArgs(args)
	if containsUnsupported, errmsg := utils.ContainsUnsupportedCharacters(input); containsUnsupported {
		log.Fatalf("[error]\n\t%s\n", errmsg)
	} else {
		println(errmsg)
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
	if asciiMap != nil {
		// everything was created fine, print...
		if flag == "justify" {
			printArt.PrintJustify(input, asciiMap)
		} else {
			printArt.PrintAlign(input, flag, asciiMap)
		}
	}
}
