package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var input string
	fmt.Printf("Type a string and I'll check if it starts with `i`, ends with `n`, and has an `a` between these characters: \n Your string: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	if strings.HasPrefix(strings.ToLower(input), "i") &&
		strings.HasSuffix(strings.ToLower(input), "n") &&
		strings.Contains(strings.ToLower(input[1:len(input)-1]), "a") {
		fmt.Printf("Found!\n")
	} else {
		fmt.Printf("Not Found!\n")
	}
}
