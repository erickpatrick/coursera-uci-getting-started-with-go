package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var name string
	var address string
	var user = make(map[string]string)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("What is your name?")
	scanner.Scan()
	name = scanner.Text()

	fmt.Println("What is your address?")
	scanner.Scan()
	address = scanner.Text()

	user["firstname"] = name
	user["address"] = address
	userJson, _ := json.Marshal(user)

	fmt.Println(string(userJson))
}
