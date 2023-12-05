package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func reverseString(str string) string {
	var output string
	for _, r := range str {
		output = string(r) + output
	}

	return output
}

func convertToDigits(str string, reversed bool) string {

	stringedDigits := map[string]string{
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

	var foundIndex int = len(str) + 1
	var foundStringedDigit string = ""

	for k := range stringedDigits {
		if reversed {
			k = reverseString(k)
		}

		wasFound := strings.Index(str, k)
		if wasFound != -1 && wasFound < foundIndex {
			foundIndex = wasFound
			foundStringedDigit = k
		}
	}

	if foundIndex <= len(str) {
		if !reversed {
			return convertToDigits(strings.Replace(str, foundStringedDigit, stringedDigits[foundStringedDigit], 1), reversed)
		} else {
			return convertToDigits(strings.Replace(str, foundStringedDigit, stringedDigits[reverseString(foundStringedDigit)], 1), reversed)
		}
	}

	return str
}

func findFirstDigit(str string) string {
	for _, r := range str {
		if unicode.IsDigit(r) {
			return string(r)
		}
	}

	return ""
}

func findLastDigit(str string) string {
	return findFirstDigit(reverseString(str))
}

func main() {
	log.SetPrefix("[advent-of-code] [day 1] [part 2] ")

	start := time.Now()

	f, err := os.Open("./input.txt")
	checkError(err)
	defer f.Close()

	var sum int = 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		first_digit := findFirstDigit(convertToDigits(line, false))
		last_digit := findLastDigit(reverseString(convertToDigits(reverseString(line), true)))

		if first_digit == "" {
			continue
		}

		converted, err := strconv.Atoi(first_digit + last_digit)
		checkError(err)

		sum += converted
	}

	elapsed := time.Since(start)

	log.Printf("Sum: %v", sum)
	log.Printf("Elapsed time: %v", elapsed)
}

func checkError(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
