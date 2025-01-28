package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var input string
	slice := make([]int, 0, 3)

	for {
		fmt.Println("Please type a valid number or 'X' to exit: ")
		fmt.Scan(&input)

		if strings.ToLower(input) == "x" {
			break
		}

		// we only want integers in the slice, so we send users
		// back to start of the loop if invalid data is entered
		number, err := strconv.Atoi(input)
		if err != nil {
			continue
		}

		slice = append(slice, number)
		sort.Ints(slice)

		fmt.Println(slice)
	}
}
