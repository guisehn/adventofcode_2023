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

var re = regexp.MustCompile(`[^0-9]`)

func keepOnlyDigits(str string) string {
	return re.ReplaceAllString(str, "")
}

func main() {
	input := readInput()
	lines := strings.Split(input, "\n")
	sum := 0

	for _, line := range lines {
		digits := keepOnlyDigits(line)
		firstAndLast := digits[0:1] + digits[len(digits)-1:]
		n, err := strconv.Atoi(firstAndLast)
		if err != nil {
			panic(err)
		}
		sum += n
	}

	fmt.Println(sum)
}
