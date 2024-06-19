package main

import (
	"log"
	"os"

	printArt "ascii-art-justify/print"
	"ascii-art-justify/utils"
)

func main() {
	// get banner file
	// turn banner file to map
	// start by opening file
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
		file, _ = os.Open(bannerfile + ".txt")
		// log.Fatal(err)
	}
	defer file.Close()

	asciiMap := utils.CreateMap(file)
	if flag == "justify" {
		printArt.PrintJustify(input, asciiMap)
	} else {
		printArt.PrintAlign(input, flag, asciiMap)
	}
}
