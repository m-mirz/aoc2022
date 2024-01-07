package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	signScore = map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
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
		playerInput := strings.Split(scanner.Text(), " ")
		fmt.Printf("%v\n", playerInput)
		score := getScore(playerInput[0], playerInput[1])
		fmt.Printf("%v\n", score)
		sum += score
	}

	fmt.Printf("Total score is: %v\n", sum)
}

func haveWon(elf string, me string) int {
	switch {
	case elf == "A" && me == "Y" || elf == "B" && me == "Z" || elf == "C" && me == "X":
		return 1 // won
	case elf == "A" && me == "X" || elf == "B" && me == "Y" || elf == "C" && me == "Z":
		return 0 // draw
	}
	return -1 // lost
}

func getScore(elf string, me string) int {
	score := 0
	won := haveWon(elf, me)
	switch won {
	case 0:
		score += 3
	case 1:
		score += 6
	}

	score += signScore[me]

	return score
}
