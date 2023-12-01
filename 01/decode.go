package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func NumberWords(s string) string {
	if strings.HasPrefix(s, "one") {
		return "1"
	}
	if strings.HasPrefix(s, "two") {
		return "2"
	}
	if strings.HasPrefix(s, "three") {
		return "3"
	}
	if strings.HasPrefix(s, "four") {
		return "4"
	}
	if strings.HasPrefix(s, "five") {
		return "5"
	}
	if strings.HasPrefix(s, "six") {
		return "6"
	}
	if strings.HasPrefix(s, "seven") {
		return "7"
	}
	if strings.HasPrefix(s, "eight") {
		return "8"
	}
	if strings.HasPrefix(s, "nine") {
		return "9"
	}
	return ""
}

func ExtractCoordinate(s string) int {
	first := ""
	last := ""
	for i, r := range s {
		if unicode.IsDigit(r) {
			if first == "" {
				first = string(s[i])
			}
			last = string(s[i])
		}
		wordDigit := NumberWords(string(s[i:]))
		if wordDigit != "" {
			if first == "" {
				first = wordDigit
			}
			last = wordDigit
		}
	}
	coordinate, err := strconv.Atoi(fmt.Sprintf("%s%s", first, last))
	if err != nil {
		return 0
	}
	return coordinate
}

func main() {

	calibrationSum := 0

	content, error := os.Open("input.txt")

	// Check whether the 'error' is nil or not. If it
	//is not nil, then print the error and exit.
	if error != nil {

		log.Fatal(error)
	}

	reader := bufio.NewScanner(content)

	for reader.Scan() {
		thisCoodinate := ExtractCoordinate(reader.Text())
		fmt.Println(strconv.Itoa(thisCoodinate) + ": " + reader.Text())
		calibrationSum += ExtractCoordinate(reader.Text())
	}
	//Print the string str
	fmt.Println(calibrationSum)
}
