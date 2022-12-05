package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type ElfGroup struct {
	assignedElves            []Elf
	doesGroupEntirelyOverlap bool
	doesGroupHaveAnyOverlap  bool
}

type Elf struct {
	startArea int
	endArea   int
}

func splitStringsToTwoIntegers(input string) (startVal int, endVal int) {
	splitStrings := strings.Split(input, "-")
	startVal, _ = strconv.Atoi(splitStrings[0])
	endVal, _ = strconv.Atoi(splitStrings[1])

	return startVal, endVal
}

func doesGroupEntirelyOverlap(group ElfGroup) (doesEntirelyOverlap bool) {
	elf1Start := group.assignedElves[0].startArea
	elf1End := group.assignedElves[0].endArea
	elf2Start := group.assignedElves[1].startArea
	elf2End := group.assignedElves[1].endArea

	if ((elf1Start <= elf2Start) && (elf1End >= elf2End)) || ((elf2Start <= elf1Start) && (elf2End >= elf1End)) {
		doesEntirelyOverlap = true
	}

	return doesEntirelyOverlap
}

func doesGroupHaveAnyOverlap(group ElfGroup) (isAnyOverlap bool) {
	elf1Start := group.assignedElves[0].startArea
	elf1End := group.assignedElves[0].endArea
	elf2Start := group.assignedElves[1].startArea
	elf2End := group.assignedElves[1].endArea

	if (elf1Start >= elf2Start && elf1Start <= elf2End) || (elf2Start >= elf1Start && elf2Start <= elf1End) {
		isAnyOverlap = true
	}
	return isAnyOverlap
}

func processElfPairs(rawLine string) (group ElfGroup) {
	splitStrings := strings.Split(rawLine, ",")
	elf1Start, elf1End := splitStringsToTwoIntegers(splitStrings[0])
	elf2Start, elf2End := splitStringsToTwoIntegers(splitStrings[1])

	elf1 := Elf{
		startArea: elf1Start,
		endArea:   elf1End,
	}

	elf2 := Elf{
		startArea: elf2Start,
		endArea:   elf2End,
	}

	group.assignedElves = append(group.assignedElves, elf1)
	group.assignedElves = append(group.assignedElves, elf2)
	group.doesGroupEntirelyOverlap = doesGroupEntirelyOverlap(group)
	group.doesGroupHaveAnyOverlap = doesGroupHaveAnyOverlap(group)

	return group
}

func parseInputFile() (groups []ElfGroup) {
	f, err := os.Open("./2022/Day4/input.txt")
	var pairs []ElfGroup

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		group := processElfPairs(line)
		pairs = append(pairs, group)
	}

	groups = pairs
	return groups
}

func main() {
	groups := parseInputFile()
	totalWithCompleteOverlap := 0
	totalWithAnyOverlap := 0

	for _, group := range groups {
		if group.doesGroupEntirelyOverlap {
			totalWithCompleteOverlap++
		}
		if group.doesGroupHaveAnyOverlap {
			totalWithAnyOverlap++
		}
	}

	fmt.Println("Total groups with complete overlap: ", totalWithCompleteOverlap)
	fmt.Println("Total groups with any overlap: ", totalWithAnyOverlap)
}
