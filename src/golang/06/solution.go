package main

import (
	"example/hello/src/golang/aocutils"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

const inputFile = "input.txt"

/* overall this puzzle can be seen as a quudratic equation
eg. test-input
time: 7
distance/record: 9

x = time the button is pressed
time to run: 7-x
overall distance covered is:
time_button_pressed * time_to_run = x(7-x) = -x^2 +7x
to tie the record the equation would be: -x^2 +7x = 9 <=> -x^2 +7x -9 = 0 <=> x^2 -7x +9 = 0
this equation can be solved and the results can be rounded to the higher and lower integer respectively

no need to algorithmic solution
pure math 🤣
*/

func solveQuadraticEquation(p float64, q float64) (float64, float64) {

	sol1 := -p/2 + math.Sqrt(math.Pow(p/2, 2.0)-q)
	sol2 := -p/2 - math.Sqrt(math.Pow(p/2, 2.0)-q)

	if sol1 < sol2 {
		return sol2, sol1
	}
	return sol1, sol2
}

func part1() {

	inputs := aocutils.ReadInput(inputFile)
	r, _ := regexp.Compile(`\d+`)
	gameTimes := r.FindAllString(inputs[0], -1)
	records := r.FindAllString(inputs[1], -1)

	solution := 1

	for game, time := range gameTimes {
		p, _ := strconv.Atoi(time)
		q, _ := strconv.Atoi(records[game])

		solHigh, solLow := solveQuadraticEquation(float64(-p), float64(q))
		/*
			button press needs to be lower then the ceil (round up) of the High value
			button press needs to be lower then the floor (round down) of the Low value
		*/
		max, min := int(math.Ceil(solHigh))-1, int(math.Floor(solLow))+1
		solution *= (max - min + 1)
	}
	fmt.Printf("[PART 1] %d\n", solution)

}

func part2() {
	inputs := aocutils.ReadInput(inputFile)
	r, _ := regexp.Compile(`\d+`)
	time := strings.Join(r.FindAllString(inputs[0], -1), "")
	record := strings.Join(r.FindAllString(inputs[1], -1), "")

	p, _ := strconv.Atoi(time)
	q, _ := strconv.Atoi(record)

	solHigh, solLow := solveQuadraticEquation(float64(-p), float64(q))
	max, min := int(math.Ceil(solHigh))-1, int(math.Floor(solLow))+1
	fmt.Printf("[PART 2] %d\n", max-min+1)
}

func main() {
	part1()
	part2()
}
