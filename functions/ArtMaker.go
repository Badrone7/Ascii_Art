package functions

import (
	"fmt"
	"os"
	"strings"
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

// the function that will check if the text contains only new lines
func Onlynewlines(text string) int {
	count := 0
	runes := []rune(text)
	if len(text) == 1 {
		return -1
	}
	for i := 0; i < len(runes); i++ {
		if i+1 < len(runes) && string(runes[i:i+2]) != "\\n" {
			return -1
		}
		count++
		if i+2 >= len(runes) {
			break
		}
		runes = runes[i+2:]
		i = -1
	}
	return count
}

// the funtion that makes the art
func ArtMaker(text string, artType int) {
	if text == "" {
		fmt.Println("Error: No text provided.")
		os.Exit(1)
	}
	if Onlynewlines(text) != -1 {
		for i := 0; i < Onlynewlines(text); i++ {
			fmt.Println()
		}
		return
	}
	art := ArtSelect(artType)
	for _, char := range text {
		if char < 32 || char > 126 {
			fmt.Println("Error: Unsupported character detected.")
			os.Exit(1)
		}
	}
	txt := strings.Split(text, "\\n")
	for i := 0; i < len(txt); i++ {
		if txt[i] == "" {
			fmt.Println()
		} else {
			PrintArt(txt[i], art)
			if i+1 < len(txt) && txt[i+1] == "\n" {
				fmt.Println()
			}
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
		if line != 8 {
			fmt.Println()
		}
	}
}

// a function that selects the art type
func ArtSelect(artType int) [][]string {
	art := [][]string{}
	switch artType {
	case 1:
		file, err := os.ReadFile("resources/shadow.txt")
		if err != nil {
			fmt.Println("Error reading file:", err)
			os.Exit(1)
		}
		art = ArtGenerator(file)
	case 2:
		file, err := os.ReadFile("resources/thinkertoy.txt")
		if err != nil {
			fmt.Println("Error reading file:", err)
			os.Exit(1)
		}
		file = []byte(strings.ReplaceAll(string(file), "\r\n", "\n"))
		art = ArtGenerator(file)
	case 0:
		file, err := os.ReadFile("resources/standard.txt")
		if err != nil {
			fmt.Println("Error reading file:", err)
			os.Exit(1)
		}
		art = ArtGenerator(file)
	}
	return art
}

func ArtGenerator(file []byte) [][]string {
	if file[0] == '\n' {
		file = file[1:]
	}
	if file[len(file)-1] != '\n' {
		file = append(file, '\n')
	}
	count := 0
	character := []string{}
	lines := []rune{}
	art := [][]string{}
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
	return art
}
