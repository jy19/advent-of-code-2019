package main

import (
	"os"
	"io/ioutil"
	"log"
	"strings"
	"strconv"
	"fmt"
)

const pwLength = 6

func findValidPasswords(input string) int {
	inputRange := strings.Split(input, "-")
	numValidPasswords := 0
	if len(inputRange) > 1 {
		start, err := strconv.Atoi(inputRange[0])
		end, err := strconv.Atoi(inputRange[1])
		if err != nil {
			log.Fatalf("couldn't convert input properly, failing..")
		}
		for i := start; i < end+1; i++ {
			if isValidPassword(strconv.Itoa(i)) {
				numValidPasswords++
			}
		}
	}

	if len(inputRange) == 1 {
		if isValidPassword(inputRange[0]) {
			return 1
		}
	}

	return numValidPasswords
}

func isValidPassword(pw string) bool {
	if len(pw) != pwLength {
		//log.Printf("length of input pw %v is not equal to pwLength %v", len(pw), pwLength)
		return false
	}
	hasRepeat := false
	noDecrease := true
	var prevChar int32
	for _, char := range pw {
		if prevChar == char {
			hasRepeat = true
		}
		if char < prevChar {
			//log.Printf("char decreased, curr %c, prev %c", char, prevChar)
			noDecrease = false
		}
		prevChar = char
	}

	return hasRepeat && noDecrease
}

func readInput(path string) (string, error) {
	byteSlice, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("error trying to read file at %v", path)
		return "", err
	}

	return string(byteSlice), nil

}

func main() {
	path := os.Args[1]

	input, err := readInput(path)
	if err != nil {
		log.Fatal(err)
	}

	numValidPw := findValidPasswords(strings.TrimSpace(input))
	fmt.Println("number of valid passwords: ", numValidPw)
}
