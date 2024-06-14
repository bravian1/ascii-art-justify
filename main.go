package main

import (
	"log"
	"os"
)

func main() {
	// get banner file
	// turn banner file to map
	// start by opening file

	bannerfile, flag, input := validateArgs(os.Args)
	if containsUnsupported, errmsg := containsUnsupportedCharacters(input); containsUnsupported {
		log.Fatalf("[error]\n\t%s\n", errmsg)
	} else {
		println(errmsg)
	}
	// fmt.Println(bannerfile)
	// fmt.Println(flag)
	// fmt.Println(input)
	file, err := os.Open(bannerfile + ".txt")
	if err != nil {
		GetFile(bannerfile + ".txt")
		file, _ = os.Open(bannerfile + ".txt")
		// log.Fatal(err)
	}
	defer file.Close()

	asciiMap := createMap(file)
	if flag == "justify" {
		printJustify(input, asciiMap)
	} else {
		printAlign(input, flag, asciiMap)
	}
}
