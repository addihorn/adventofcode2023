package main

import (
	"example/hello/src/golang/aocutils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const inputFile = "input.txt"

var maxCubes = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func part2() {

	inputData := aocutils.ReadInput(inputFile)
	sum := 0

	for _, input := range inputData {
		//fmt.Println(input)
		gameData := strings.Split(input, ":")
		gameTurns := gameData[1]
		gamePower := 1

		var minCubes = map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for colour, min := range minCubes {
			r, _ := regexp.Compile(fmt.Sprintf(`\d+ %s`, colour))
			matches := r.FindAllString(gameTurns, -1)

			for _, match := range matches {
				//fmt.Fprintln(os.Stderr, strings.Split(match, " ")[0])
				cubeCount, _ := strconv.Atoi(strings.Split(match, " ")[0])
				//fmt.Printf("Comparing %s: %d min %d \n", colour, cubeCount, min)
				if cubeCount > min {
					min = cubeCount
				}
			}
			minCubes[colour] = min
			gamePower *= min
		}

		sum += gamePower
	}

	fmt.Printf("=====> Solution Part 2: %d <=====\n", sum)
}

func part1() {

	inputData := aocutils.ReadInput(inputFile)
	sum := 0

	for _, input := range inputData {
		//fmt.Println(input)
		gameData := strings.Split(input, ":")
		gameNumber := gameData[0]
		gameTurns := gameData[1]
		impossible := false

		for colour, max := range maxCubes {
			r, _ := regexp.Compile(fmt.Sprintf(`\d+ %s`, colour))
			matches := r.FindAllString(gameTurns, -1)

			for _, match := range matches {
				//fmt.Fprintln(os.Stderr, strings.Split(match, " ")[0])
				cubeCount, _ := strconv.Atoi(strings.Split(match, " ")[0])
				//fmt.Printf("Comparing %s: %d max %d \n", colour, cubeCount, max)
				if cubeCount > max {
					impossible = true
				}

			}
		}

		if !impossible {
			gameNumber = strings.ReplaceAll(gameNumber, "Game ", "")
			gameNo, _ := strconv.Atoi(gameNumber)
			sum += gameNo
		}

	}

	fmt.Printf("=====> Solution Part 1: %d <=====\n", sum)
}

func main() {
	part1()

	fmt.Println(fmt.Sprintln())

	part2()
}
