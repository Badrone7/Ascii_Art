package functions

import (
	"fmt"
)

// a function that checks the arguments
func ArgsChecker(args []string) bool {
	if len(args) < 2 {
		fmt.Println("Not enough arguments")
		return false
	} else if len(args) > 2 {
		fmt.Println("Too many arguments")
		return false
	}
	return true
}
