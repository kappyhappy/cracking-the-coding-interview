package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isUniqueChars2(str string) (bool, string) {
	if len(str) > 256 {
		return false, "please use ascii chars"
	}

	var char_set [256]bool

	for i := 0; i < len(str); i++ {
		var val int = int(str[i])
		if char_set[val] {
			return false, "'" + str + "' contains same char '" + string(rune(str[i])) + "'"
		}
		char_set[val] = true
	}
	return true, str + " does not use duplicated chars"
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
	isUniqueChar, message := isUniqueChars2(stringInput)

	fmt.Println("check result is...")
	fmt.Println("input consists of unique char? " + strconv.FormatBool(isUniqueChar))
	fmt.Println(message)
}
