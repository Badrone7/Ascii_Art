package functions

import (
	"fmt"
	"os"
	"strings"
)

// the function that will check if the text contains only new lines
func Onlynewlines(text string) int {
	count := 0
	runes := []rune(text)
	for i := 0; i < len(runes); i++ {
		if runes[i] != '\\' {
			return -1
		}
		if i+1 >= len(runes) || runes[i+1] != 'n' {
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
func ArtMaker(text string) {
	if text == "" {
		return
	}
	if Onlynewlines(text) != -1 {
		for i := 0; i < Onlynewlines(text); i++ {
			fmt.Println()
		}
		return
	}
	art := ArtSelect()
	for _, char := range text {
		if char < 32 || char > 126 {
			fmt.Println("Error: Unsupported character detected.")
			return
		}
	}
	txt := strings.Split(string(text), "\\n")
	for i := 0; i < len(txt); i++ {
		if txt[i] == "" {
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
		if line != 8 {
			fmt.Println()
		}
	}
}

// a function that selects the art type
func ArtSelect() [][]string {
	file, err := os.ReadFile("resources/standard.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
	art := ArtGenerator(file)
	return art
}

// a function that generates the art from a txt file
func ArtGenerator(file []byte) [][]string {
	if file[0] == '\n' {
		file = file[1:]
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
