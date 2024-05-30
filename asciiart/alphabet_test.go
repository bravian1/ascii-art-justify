package asciiart

import (
	"os"
	"path/filepath"
	"testing"
)

/*
* TestAlphabet(t *testing.T)
*
* Tests the function createAlphabet in main.go
* This file tests for:
*	1) the fact that given a valid file, the alphabet created will be non-empty
*	2) given a letter that should be present in the map, the corresponding value in the map is not empty
 */
func TestAlphabet(t *testing.T) {
	path := filepath.Join("..", "standard.txt")
	file, err := os.Open(path)
	if err != nil {
		t.Fatalf("Failed to open: %v", path)
	}
	defer file.Close()
	// successfully created the map
	if ok, _ := CreateAlphabet(path); !ok {
		t.Errorf("Expected a map, but got nil instead\n")
	}
}

func TestNonExistent(t *testing.T) {
	path := filepath.Join("..", "nonexistent.txt")
	file, err := os.Open(path)
	if err == nil {
		t.Errorf("Should not be able to open non existent file %v", path)
	}
	defer file.Close()
	// expect the map to be nil
	// a non existent file
	if ok, _ := CreateAlphabet(path); ok {
		t.Errorf("Expected nil, but got a map instead\n")
	}
}

/*
*	Test for the presence or lack thereof a few letters within the map
*
 */
func TestAFewLetters(t *testing.T) {
	type Letter struct {
		letter rune
		expect bool
	}
	table := []Letter{
		{'h', true},
		{'A', true},
		{'\n', false},
		{'8', true},
		{'\r', false},
		{'\b', false},
		{'\\', true},
	}

	path := filepath.Join("..", "standard.txt")
	file, err := os.Open(path)
	if err != nil {
		t.Fatalf("Failed to open: %v", path)
	}
	defer file.Close()
	var ok bool
	var alphabet map[rune][]string
	if ok, alphabet = CreateAlphabet(path); ok {
		for _, item := range table {
			if _, present := alphabet[item.letter]; present != item.expect {
				t.Errorf("For '%c', expected '%v', but got '%v' instead\n", item.letter, item.expect, present)
			}
		}
	}
}
