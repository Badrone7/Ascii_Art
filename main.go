package main

import (
	"os"

	"ascii/functions"
)

/*
	In this Program we gonna turn a simple ascii text into an ASCII Art
	The program will take 2 arguments:
		- the file containing the main function
		- The text to convert
	Example of usage:
		go run main.go "Hello World"
*/

// the main function

func main() {
	validarg := functions.ArgsChecker(os.Args)
	if !validarg {
		return
	}
	functions.ArtMaker(os.Args[1])
}
