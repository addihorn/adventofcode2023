package main

import (
	"example/hello/src/golang/aocutils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const inputFile = "input.txt"

func part1() {
	inputData := aocutils.ReadInput(inputFile)
	r, _ := regexp.Compile(`\d`)
	sum := int64(0)
	for _, line := range inputData {
		fmt.Println(line)
		matches := r.FindAllString(line, -1)
		number := fmt.Sprintf("%s%s", matches[0], matches[len(matches)-1])
		i, _ := strconv.ParseInt(number, 10, 0)
		fmt.Printf("%s %d", number, i)
		sum = sum + i
		fmt.Println(sum)

	}

}

func part2() {
	numberMap := map[string]string{
		"two":   "2",
		"one":   "1",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	inputData := aocutils.ReadInput(inputFile)

	r, _ := regexp.Compile(`\d`)
	sum := int64(0)
	for _, line := range inputData {
		for key, value := range numberMap {
			line = strings.ReplaceAll(line, key, fmt.Sprintf("%s%s%s", key, value, key))
		}

		fmt.Println(line)
		matches := r.FindAllString(line, -1)
		number := fmt.Sprintf("%s%s", matches[0], matches[len(matches)-1])
		i, _ := strconv.ParseInt(number, 10, 0)
		fmt.Printf("%s %d", number, i)
		sum = sum + i
		fmt.Println(sum)

	}

	fmt.Println(sum)

}

func main() {
	part2()
}
