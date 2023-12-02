package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readInput() string {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(dat))
}

var wordToDigits = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

var wordsRe = regexp.MustCompile("(one|two|three|four|five|six|seven|eight|nine)")

func replaceWords(str string) string {
	for {
		prevStr := str

		str = wordsRe.ReplaceAllStringFunc(str, func(m string) string {
			return wordToDigits[m] + m[len(m)-1:]
		})

		if str == prevStr {
			break
		}
	}

	return str
}

var nonDigitsRe = regexp.MustCompile(`[^0-9]`)

func keepOnlyDigits(str string) string {
	return nonDigitsRe.ReplaceAllString(str, "")
}

func main() {
	input := readInput()
	lines := strings.Split(input, "\n")
	sum := 0

	for _, line := range lines {
		digits := keepOnlyDigits(replaceWords(line))
		fmt.Println(digits)
		fmt.Println("---")

		firstAndLast := digits[0:1] + digits[len(digits)-1:]
		n, err := strconv.Atoi(firstAndLast)
		if err != nil {
			panic(err)
		}
		sum += n
	}

	fmt.Println(sum)
}
