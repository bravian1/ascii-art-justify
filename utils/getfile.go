package utils

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"time"
)

func GetFile(fileName string) {
	// check for internet connection
	if _, err := net.DialTimeout("tcp", "www.google.com:443", time.Second*5); err != nil {
		fmt.Println("Error: no internet connection")
		return
	}

	fullURL := "https://learn.zone01kisumu.ke/git/root/public/raw/branch/master/subjects/ascii-art/" + fileName

	// create a blank file
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error:", err)
	}
	// create a client
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			req.URL.Opaque = req.URL.Path
			return nil
		},
	}

	// Put fetched contents on file
	response, err := client.Get(fullURL)
	if err != nil {
		// timeout error
		if networkError, ok := err.(net.Error); ok && networkError.Timeout() {
			fmt.Println("Error: Timeout occurred while connecting to the internet.")
		} else {
			fmt.Println("Error:", err)
		}
	}

	defer response.Body.Close()
	size, err := io.Copy(file, response.Body)
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer file.Close()
	fmt.Printf("Downloaded a file %s with size %d\n", fileName, size)
}
