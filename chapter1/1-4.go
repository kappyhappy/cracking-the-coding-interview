package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func replaceSpaces(str string, length int) string {
	spaceCount := 0
	newLength := 0
	rune_array := []rune(str)

	for i := 0; i < length; i++ {
		if string(rune_array[i]) == " " {
			spaceCount++
		}
	}

	newLength = length + spaceCount*2
	if newLength == length {
		return str
	}

	new_rune_array := []rune(strings.Repeat(" ", newLength))

	for i := length - 1; i >= 0; i-- {
		if string(rune_array[i]) == " " {
			new_rune_array[newLength-1] = []rune("0")[0]
			new_rune_array[newLength-2] = []rune("2")[0]
			new_rune_array[newLength-3] = []rune("%")[0]
			newLength = newLength - 3
		} else {
			new_rune_array[newLength-1] = rune_array[i]
			newLength = newLength - 1
		}
	}

	return string(new_rune_array)
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
	stringInput := stdinToString()
	message := replaceSpaces(stringInput, len([]rune(stringInput)))

	fmt.Println("check result is...")
	fmt.Println("input strings is: " + stringInput)
	fmt.Println("replaced strings from ' ' to '%20' is: " + message)
}
