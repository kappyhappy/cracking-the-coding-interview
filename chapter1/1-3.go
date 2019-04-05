package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func permutation(str1 string, str2 string) (bool, string) {
	if len(str1) != len(str2) {
		return false, "length for 2 strings does not match"
	}

	var letters [256]int // ascii is assumed

	for i := 0; i < len(str1); i++ {
		letters[str1[i]]++
	}

	for i := 0; i < len(str2); i++ {
		var c byte = str2[i]
		letters[c]--
		if letters[c] < 0 {
			return false, "2 strings are not anagrams"
		}
	}
	return true, "2 strings are anagrams"
}

func stdinToString() (stringInput string) {
	fmt.Println("type the word:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var stdinString = scanner.Text()
	stdinString = strings.TrimSpace(stdinString)
	return stdinString
}

func main() {
	stringInput1 := stdinToString()
	stringInput2 := stdinToString()
	isAnagrams, message := permutation(stringInput1, stringInput2)

	fmt.Println("check result is...")
	fmt.Println("input strings are anagrams? " + strconv.FormatBool(isAnagrams))
	fmt.Println(message)
}
