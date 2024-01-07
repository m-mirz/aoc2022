package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	shapeScore = map[string]int{
		"A": 1, // rock
		"B": 2, // paper
		"C": 3, // scissors
	}

	shapeRules = map[string][2]string{
		// winning, losing shape
		"A": {"B", "C"},
		"B": {"C", "A"},
		"C": {"A", "B"},
	}

	resultScore = map[string]int{
		"X": 0, // lose
		"Y": 3, // draw
		"Z": 6, // win
	}
)

func main() {
	readFile, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	sum := 0
	for scanner.Scan() {
		lineInput := strings.Split(scanner.Text(), " ")
		fmt.Printf("%v\n", lineInput)
		score := getScore(lineInput[0], lineInput[1])
		fmt.Printf("%v\n", score)
		sum += score
	}

	fmt.Printf("Total score is: %v\n", sum)
}

func getScore(elf string, haveWon string) int {
	score := resultScore[haveWon]
	switch haveWon {
	case "X": // lose
		score += shapeScore[shapeRules[elf][1]]
	case "Y": // draw
		score += shapeScore[elf]
	case "Z": // win
		score += shapeScore[shapeRules[elf][0]]
	}

	return score
}
