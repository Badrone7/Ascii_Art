package main

import (
	"ascii/functions"
	"os"
)

/*
	In this Program we gonna turn a simple ascii text into an ASCII Art
	We have 3 types of arts:
		- Standard
		- Shadow
		- Thinkertoy
	The program will take 2 arguments:
		- The text to convert
		- The command (optional)
	Example of usage:
		go run main.go "Hello World" --shadow
*/

// the main function

func main() {
	validarg, artType := functions.ArgsChecker(os.Args)
	if !validarg {
		return
	}
	functions.ArtMaker(os.Args[1], artType)
}
