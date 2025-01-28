package main

import "fmt"

func main() {
	var userNumber float32

	fmt.Printf("Type a number: ")
	fmt.Scan(&userNumber)
	fmt.Printf("Your number is %d\n", int64(userNumber))
}
