package functions

import (
	"fmt"
	"os"
)

// a functions that splits the text by new lines
func Splitter(runes []rune) []string {
	parts := []string{}
	for i := 0; i < len(runes); i++ {
		if i+1 < len(runes) && runes[i] == '\\' {
			if runes[i+1] == 'n' {
				parts = append(parts, string(runes[:i]))
				parts = append(parts, "\n")
				if i+2 >= len(runes) {
					return parts
				}
				runes = runes[i+2:]
				i = -1
			}
		}
	}
	if len(runes) > 0 {
		parts = append(parts, string(runes))
	}
	return parts
}

// the funtion that makes the art
func ArtMaker(text string, artType int) {
	if text == "" {
		fmt.Println("Error: No text provided.")
		os.Exit(1)
	}
	if text == "\\n" {
		fmt.Println()
		return
	}
	art := ArtSelect(artType)
	for _, char := range text {
		if char < 32 || char > 126 {
			fmt.Println("Error: Unsupported character detected.")
			os.Exit(1)
		}
	}
	runes := []rune(text)
	txt := Splitter(runes)
	for i := 0; i < len(txt); i++ {
		if txt[i] == "\n" {
			fmt.Println()
		} else {
			PrintArt(txt[i], art)
		}
	}
}

// a function that prints the art
func PrintArt(text string, art [][]string) {
	for line := 0; line < 8; line++ {
		for _, char := range text {
			index := int(char) - 32
			fmt.Print(art[index][line])
		}
		fmt.Println()
	}
}

// a function that selects the art type
func ArtSelect(artType int) [][]string {
	art := [][]string{}
	character := []string{}
	lines := []rune{}
	if artType == 1 {
		count := 0
		file, err := os.ReadFile("resources/shadow.txt")
		if err != nil {
			fmt.Println("Error reading file:", err)
			os.Exit(1)
		}
		if file[0] == '\n' {
			file = file[1:]
		}
		for _, c := range string(file) {
			if count == 8 {
				art = append(art, character)
				count = 0
				character = []string{}
				continue
			}
			if c == '\n' {
				count++
				character = append(character, string(lines))
				lines = []rune{}
				continue
			}
			lines = append(lines, c)
		}
		if len(character) > 0 {
			art = append(art, character)
		}
	} else if artType == 2 {
		count := 0
		file, err := os.ReadFile("resources/thinkertoy.txt")
		if err != nil {
			fmt.Println("Error reading file:", err)
			os.Exit(1)
		}
		if file[0] == '\n' {
			file = file[1:]
		}
		for _, c := range string(file) {
			if count == 8 {
				art = append(art, character)
				count = 0
				character = []string{}
				continue
			}
			if c == '\n' {
				count++
				character = append(character, string(lines))
				lines = []rune{}
				continue
			}
			lines = append(lines, c)
		}
		if len(character) > 0 {
			art = append(art, character)
		}
	} else if artType == 0 {
		count := 0
		file, err := os.ReadFile("resources/standard.txt")
		if err != nil {
			fmt.Println("Error reading file:", err)
			os.Exit(1)
		}
		if file[0] == '\n' {
			file = file[1:]
		}
		for _, c := range string(file) {
			if count == 8 {
				art = append(art, character)
				count = 0
				character = []string{}
				continue
			}
			if c == '\n' {
				count++
				character = append(character, string(lines))
				lines = []rune{}
				continue
			}
			lines = append(lines, c)
		}
		if len(character) > 0 {
			art = append(art, character)
		}
	}
	return art
}
