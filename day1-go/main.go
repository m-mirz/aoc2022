package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {

	elves := make([]int, 0)

	readFile, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	cals_sum := 0
	for fileScanner.Scan() {
		if fileScanner.Text() == "" {
			// add elf
			elves = append(elves, cals_sum)
			cals_sum = 0
			continue
		}

		cals, err := strconv.Atoi(fileScanner.Text())
		if err != nil {
			fmt.Println("Error during conversion")
			return
		}

		cals_sum += cals
	}

	readFile.Close()

	max_elf := 0
	max_elf_idx := 0

	for i, cals := range elves {
		if cals > max_elf {
			max_elf = cals
			max_elf_idx = i
		}
	}

	fmt.Println("Elf " + strconv.Itoa(max_elf_idx) + " has " + strconv.Itoa(max_elf) + " calories")

	sort.Slice(elves, func(i, j int) bool {
		return elves[i] < elves[j]
	})

	fmt.Println(elves)

}
