package functions

import (
	"fmt"
	"os"
	"regexp"
)

// a function that checks the arguments
func ArgsChecker(args []string) (bool, int) {
	if len(args) < 2 {
		fmt.Println("Not enough arguments")
		return false, 0
	} else if len(args) > 3 {
		fmt.Println("Too many arguments")
		return false, 0
	} else if len(args) == 3 {
		re := regexp.MustCompile(`^--(help|h|howto)$`)
		if re.MatchString(args[2]) {
			fmt.Println("Usage: ascii [options]\nOptions:\n  \"--shadow\" or \"--standard\" or \"--thinkertoy\"\n\nExemple:\n  go run main.go --shadow")
			os.Exit(0)
		}
		re = regexp.MustCompile(`^--(shadow|standard|thinkertoy)$`)
		if !re.MatchString(args[2]) {
			fmt.Println("Invalid argument:", args[2])
			return false, 0
		}
	} else {
		return true, 0
	}
	return true, ArtDetect(args[2])
}

// a function that detects the art type
func ArtDetect(arg string) int {
	switch arg {
	case "--shadow":
		return 1
	case "--thinkertoy":
		return 2
	case "--standard":
		return 0
	default:
		return 0
	}
}
