package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const MAX_CUBES_RED = 12
const MAX_CUBES_GREEN = 13
const MAX_CUBES_BLUE = 14

const GAME_PREFIX_LEN = len("Game ")

func getPowersPerColor(line string) map[string]int {
	game := strings.Split(line, ": ")[1]

	cubeMaxes := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	cubeSets := strings.Split(game, "; ")
	for i := 0; i < len(cubeSets); i++ {

		cubes := strings.Split(cubeSets[i], ", ")

		for j := 0; j < len(cubes); j++ {
			cubeInfo := strings.Split(cubes[j], " ")

			cubeCount, err := strconv.Atoi(cubeInfo[0])
			checkError(err)

			if cubeCount > cubeMaxes[cubeInfo[1]] {
				cubeMaxes[cubeInfo[1]] = cubeCount
			}
		}
	}

	return cubeMaxes
}

func getPower(line string) int {
	powersPerColor := getPowersPerColor(line)

	var power int = 1

	for _, v := range powersPerColor {
		power = power * v
	}

	return power
}

func main() {
	log.SetPrefix("[advent-of-code] [day 2] [part 1] ")

	f, err := os.Open("./input.txt")
	checkError(err)
	defer f.Close()

	sum := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		sum += getPower(line)
	}

	log.Printf("Sum of powers: %v", sum)
}

func checkError(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
