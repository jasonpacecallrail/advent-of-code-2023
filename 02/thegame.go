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

type Bag struct {
	red, green, blue int
}

type Round struct {
	red, green, blue int
}

type Game struct {
	id              int
	one, two, three Round
}

func ParseGame(s string) Game {
	split = strings.Split(s, ":")
}

func GamePossible(s string, b Bag) int {
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

func RoundPossible(s string)

func main() {

	thisBag := Bag{12, 13, 14}

	roundSum := 0

	content, error := os.Open("input.txt")

	// Check whether the 'error' is nil or not. If it
	//is not nil, then print the error and exit.
	if error != nil {

		log.Fatal(error)
	}

	reader := bufio.NewScanner(content)

	for reader.Scan() {
		roundSum += GamePossible(reader.Text(), thisBag)
	}
	//Print the string str
	fmt.Println(roundSum)
}
