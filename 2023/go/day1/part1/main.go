package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"unicode"
)

func reverseString(str string) string {
	var output string
	for _, r := range str {
		output = string(r) + output
	}

	return output
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
	log.SetPrefix("[advent-of-code] [day 1] [part 1] ")

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

		first_digit := findFirstDigit(line)
		last_digit := findLastDigit(line)

		if first_digit == "" {
			continue
		}

		converted, err := strconv.Atoi(first_digit + last_digit)
		checkError(err)

		sum += converted
	}

	log.Printf("Sum: %v", sum)
}

func checkError(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
