package main

import (
	"fmt"
	"os"
	"strings"
	"ascii/functions"
)

func main() {
	args := os.Args
	// Check if the correct number of arguments is provided
	if len(args) != 2 {
		fmt.Println("Error: invalid arguments")
		os.Exit(1)
	}
	txt := args[1]

	// Split the input text by "\\n" to handle newlines
	textSlice := strings.Split(txt, "\\n")

	// Validate the characters in the input text
	if !functions.CharValidation(txt) {
		fmt.Println("Error: invalid char")
		os.Exit(1)
	}

	// Read the content of the standard.txt file
	file, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error: reading file")
		os.Exit(1)
	}
	// Split the file content by newlines to create a slice of ASCII art lines
	slice := strings.Split(string(file), "\n")

	// Generate and print the ASCII art using the input text and the ASCII art lines
	functions.Ascii(slice, textSlice)
}
