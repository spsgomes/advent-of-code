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

func isValidGame(line string) bool {
	game := strings.Split(line, ": ")[1]

	cubeSets := strings.Split(game, "; ")
	for i := 0; i < len(cubeSets); i++ {

		cubeTotals := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		cubes := strings.Split(cubeSets[i], ", ")

		for j := 0; j < len(cubes); j++ {
			cubeInfo := strings.Split(cubes[j], " ")

			cubeCount, err := strconv.Atoi(cubeInfo[0])
			checkError(err)

			cubeTotals[cubeInfo[1]] += cubeCount
		}

		if cubeTotals["red"] > MAX_CUBES_RED || cubeTotals["green"] > MAX_CUBES_GREEN || cubeTotals["blue"] > MAX_CUBES_BLUE {
			return false
		}
	}

	return true
}

func getGameId(line string) int {
	split := strings.Split(line, ":")

	line = split[0]

	gameId, err := strconv.Atoi(line[GAME_PREFIX_LEN:])
	checkError(err)

	return gameId
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

		if isValidGame(line) {
			sum += getGameId(line)
		}
	}

	log.Printf("Sum of games: %v", sum)
}

func checkError(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
