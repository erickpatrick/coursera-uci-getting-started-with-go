package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func main() {
	var nameEntries = make([]name, 0, 1)
	var filename string

	fmt.Println("what's the name of the file")
	fmt.Scan(&filename)

	readFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()

	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		nameEntries = append(nameEntries, name{fname: split[0], lname: split[1]})
	}

	for _, name := range nameEntries {
		fmt.Println(name.fname, name.lname)
	}
}
