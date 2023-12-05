package main

import (
	"example/hello/src/golang/aocutils"
	"fmt"
	"math"
	"regexp"
	"strings"
)

var inputFile = "input.txt"

type Scratchcard struct {
	iHave          int
	winningNumbers map[string]string
	scratchNumbers []string
}

func newScratchcard(data string) *Scratchcard {
	r, _ := regexp.Compile(`\d+`)
	scratchCard := strings.Split(data, ":")[1]
	numbers := strings.Split(scratchCard, "|")
	winningNumbersAsList := r.FindAllString(numbers[0], -1)
	myNumbers := r.FindAllString(numbers[1], -1)

	winningNumbers := make(map[string]string)
	for _, winningNumber := range winningNumbersAsList {
		winningNumbers[winningNumber] = winningNumber
	}

	return &Scratchcard{iHave: 1, winningNumbers: winningNumbers, scratchNumbers: myNumbers}
}

func (card *Scratchcard) getPoints() int {
	wins := 0

	for _, number := range card.scratchNumbers {
		if card.winningNumbers[number] != "" {
			wins++
		}
	}
	return int(math.Pow(2.0, float64(wins-1)))
}

func (card *Scratchcard) getMatchingNumbers() int {
	wins := 0

	for _, number := range card.scratchNumbers {
		if card.winningNumbers[number] != "" {
			wins++
		}
	}
	return wins
}

func part1() {

	inputs := aocutils.ReadInput(inputFile)
	points := 0
	for _, input := range inputs {

		scratchcard := newScratchcard(input)

		points += scratchcard.getPoints()

	}

	fmt.Printf("[PART 1] %d points \n", points)

}

func part2() {
	inputs := aocutils.ReadInput(inputFile)
	numberOfGames := 0

	//save all games in array
	uniqueScratchcards := make([]*Scratchcard, len(inputs))

	for index, input := range inputs {
		uniqueScratchcards[index] = newScratchcard(input)
	}
	for gameNumber, scratchCard := range uniqueScratchcards {
		numberOfGames += scratchCard.iHave
		additionalGames := scratchCard.getMatchingNumbers()
		//fmt.Printf("Additional games from game %d: %d \n", gameNumber, additionalGames)
		for i := 1; i < additionalGames+1; i++ {
			//fmt.Printf("Adding to game %d from game %d \n", gameNumber+i, gameNumber)
			uniqueScratchcards[gameNumber+i].iHave += scratchCard.iHave
		}
	}
	fmt.Printf("[PART 2] %d total scratchcards \n", numberOfGames)
}

func main() {
	part1()
	part2()
}
