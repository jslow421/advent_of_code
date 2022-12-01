package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"time"
)

type Elf struct {
	name          string
	foodItems     []Entry
	totalCalories int
}

func (e Elf) TotalCalories() int64 {
	var total = int64(0)

	for _, element := range e.foodItems {
		total += element.calories
	}

	return total
}

type Entry struct {
	calories int64
}

func parseInputFile() (elves []Elf) {
	f, err := os.Open("./2022/Day1/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	elfCount := 0
	elf := Elf{
		name:          strconv.FormatInt(int64(elfCount), 10),
		totalCalories: 0,
		foodItems:     []Entry{},
	}

	for scanner.Scan() {
		count, _ := strconv.ParseInt(scanner.Text(), 0, 64)
		elf.foodItems = append(elf.foodItems, Entry{calories: count})

		if scanner.Text() == "" {
			elves = append(elves, elf)
			elfCount++
			elf = Elf{
				name:          strconv.FormatInt(int64(elfCount), 10),
				totalCalories: 0,
				foodItems:     []Entry{},
			}
		}
	}

	return elves
}

func sortElvesByProduction(elves []Elf) {
	sort.Slice(elves, func(i int, j int) bool {
		return elves[i].TotalCalories() > elves[j].TotalCalories()
	})
}

func calculateTotalForGroup(elves []Elf) (totalCalories int64) {
	for _, elf := range elves {
		totalCalories += elf.TotalCalories()
	}

	return totalCalories
}

func main() {
	start := time.Now()
	elves := parseInputFile()
	sortElvesByProduction(elves)
	total := calculateTotalForGroup(elves[:3])
	for _, elf := range elves[:3] {
		fmt.Println(elf.name)
		fmt.Println(elf.TotalCalories())
		fmt.Println("")
	}
	fmt.Println("Total for top 3: ", total)
	fmt.Println(time.Since(start))
}
