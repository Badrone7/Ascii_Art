package main

import (
	"fmt"
	"os"
)

func ArtTaker(str string) []string {
	data, err := os.ReadFile(str)
	if err != nil {
		panic(err)
	}
	characters := []string{}
	runesData := []rune(string(data))
	lines := 0
	char := ""
	count := 0
	for i := 0; i < len(runesData); i++ {
		if runesData[i] == '\n' {
			lines++
		}
		if lines >= 10 {
			char += string(runesData[i])
			if runesData[i] == '\n' {
				count++
			}
			if count == 8 {
				if char != "" {
					characters = append(characters, char)
				}
				char = ""
				count = 0
			}
		}
	}
	return characters
}

func main() {
	var characters []string
	if len(os.Args) < 2 {
		fmt.Println("There is no Text to print Art for")
		return
	}
	if len(os.Args) > 3 {
		fmt.Println("Too many arguments")
		return
	}
	if len(os.Args) == 3 {
		Artarg := []rune(os.Args[2])
		if string(Artarg[:2]) == "--" {
			Artarg = Artarg[2:]
			if Artarg[0] == 'h' || Artarg[0] == '?' || string(Artarg[:4]) == "help" {
				if (Artarg[0] == 'h' || Artarg[0] == '?') && string(Artarg[:4]) != "help" {
					fmt.Println("I guess you meant \"help\"")
				}
				fmt.Println("\nHelp:\nThis program prints the Text Provided as Art\nUsage: go run \"nameofprogram.go\" \"Text to print as Art\" \"--Type of Art\"\n(The Type of Art available are : standard, shadow, thinkertoy)\nExample: go run main.go \"Hello World\" --standard\n")
				return
			}
			if string(Artarg) != "standard" && string(Artarg) != "shadow" && string(Artarg) != "thinkertoy" {
				fmt.Println("Wrong Art Type")
				return
			} else {
				characters = ArtTaker(string(Artarg) + ".txt")
			}
		}
	} else {
		characters = ArtTaker("standard.txt")
	}
	txt := []rune(os.Args[1])
	splitted := []string{}
	for i := 0; i < len(txt); i++ {
		if i+1 != len(txt)-1 && txt[i] == '\\' && txt[i+1] == 'n' {
			splitted = append(splitted, string(txt[:i]))
			txt = txt[i+2:]
			i = 0
		}
	}
	if len(txt) > 0 {
		splitted = append(splitted, string(txt))
	}
	//12
	lines := 0
	charline:=""
	charlines := []string{}
	for i:=0;i<len(characters);i++{
		for j:=0 ; j<len(characters[i]);j++{
			if characters[i][j]=='\n'{
				lines++
			}
		}
		charline+=string(characters[i])
}
