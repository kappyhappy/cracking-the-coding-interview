package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverse(str string) {
	var pointer *string = &str
	fmt.Println("Golang does not allow null string")
	fmt.Println("given string is")
	fmt.Println(str)
	fmt.Println("pointer for given string is")
	fmt.Println(pointer)
}

func stdinToString() (stringInput string) {
	fmt.Println("type the word to check if it uses same letters:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var stdinString = scanner.Text()
	stdinString = strings.TrimSpace(stdinString)
	return stdinString
}

func main() {
	stringInput := stdinToString()
	fmt.Println("check result is...")
	reverse(stringInput)
}
