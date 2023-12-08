package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func getSurroundingMatrix(matrix [][]string, i int, j int) [][]string {

	compare := make([][]string, 3)

	compare[0] = []string{
		matrix[i-1][j-1],
		matrix[i-1][j],
		matrix[i-1][j+1],
	}

	compare[1] = []string{
		matrix[i][j-1],
		".",
		matrix[i][j+1],
	}

	compare[2] = []string{
		matrix[i+1][j-1],
		matrix[i+1][j],
		matrix[i+1][j+1],
	}

	return compare
}

func hasNeighboringSymbols(m [][]string, i int, j int) bool {

	matrix := getSurroundingMatrix(m, i, j)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] != "." && !unicode.IsDigit([]rune(matrix[i][j])[0]) {
				return true
			}
		}
	}

	return false
}

func findFullNumbers(matrix [][]string, i int, j int) (int, int, int) {

	number := ""
	minj, maxj := j, j

	// Search to the back
	for j2 := j; j2 >= 0; j2-- {
		if unicode.IsDigit([]rune(matrix[i][j2])[0]) {
			number = matrix[i][j2] + number
			minj = j2
		} else {
			break
		}
	}

	// Search to the front
	for j2 := j + 1; j2 < len(matrix[i]); j2++ {
		if unicode.IsDigit([]rune(matrix[i][j2])[0]) {
			number = number + matrix[i][j2]
			maxj = j2
		} else {
			break
		}
	}

	result, err := strconv.Atoi(number)
	checkError(err)

	return result, minj, maxj
}

func convertToMatrix(originalData string) [][]string {

	data := strings.Split(originalData, "\n")
	for i := 0; i < len(data); i++ {
		data[i] = strings.TrimSpace(data[i])
	}

	// Pad the matrix by 1 on each side
	arr := make([][]string, len(data)+2, len(data[0])+2)

	for i := 0; i < len(data)+2; i++ {
		arr[0] = append(arr[0], ".")
		arr[len(arr)-1] = append(arr[0], ".")
	}

	for row := 0; row < len(data); row++ {
		splitChars := strings.Split(data[row], "")
		splitChars = append(splitChars, ".")
		splitChars = append([]string{"."}, splitChars...)

		arr[row+1] = splitChars
	}

	return arr
}

func parsePartNumbers(data string) []int {

	matrix := convertToMatrix(data)

	var arr []int

	for i := 1; i < len(matrix)-1; i++ {
		var minj, maxj int = -1, -1

		for j := 1; j < len(matrix[i])-1; j++ {
			if unicode.IsDigit([]rune(matrix[i][j])[0]) {
				if hasNeighboringSymbols(matrix, i, j) {
					fullNumber, minjTmp, maxjTmp := findFullNumbers(matrix, i, j)

					if minjTmp != minj && maxjTmp != maxj {
						arr = append(arr, fullNumber)

						minj, maxj = minjTmp, maxjTmp
					}
				}
			}
		}
	}

	return arr
}

func main() {
	log.SetPrefix("[advent-of-code] [day 3] [part 1] ")

	d, err := os.ReadFile("./input.txt")
	checkError(err)

	data := strings.TrimSpace(string(d))

	sum := arraySum(parsePartNumbers(data))

	log.Printf("Sum of games: %v", sum)
}

func arraySum(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func checkError(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
