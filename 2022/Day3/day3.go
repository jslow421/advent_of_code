package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

type Bag struct {
	firstCompartment        []Item
	secondCompartment       []Item
	itemsInBothCompartments []Item
}

type Item struct {
	value         string
	priorityScore int
}

type Group struct {
	bags       []Bag
	groupType  string
	groupScore int
}

var lowercaseMap = map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"f": 6,
	"g": 7,
	"h": 8,
	"i": 9,
	"j": 10,
	"k": 11,
	"l": 12,
	"m": 13,
	"n": 14,
	"o": 15,
	"p": 16,
	"q": 17,
	"r": 18,
	"s": 19,
	"t": 20,
	"u": 21,
	"v": 22,
	"w": 23,
	"x": 24,
	"y": 25,
	"z": 26,
}

var uppercaseMap = map[string]int{
	"A": 27,
	"B": 28,
	"C": 29,
	"D": 30,
	"E": 31,
	"F": 32,
	"G": 33,
	"H": 34,
	"I": 35,
	"J": 36,
	"K": 37,
	"L": 38,
	"M": 39,
	"N": 40,
	"O": 41,
	"P": 42,
	"Q": 43,
	"R": 44,
	"S": 45,
	"T": 46,
	"U": 47,
	"V": 48,
	"W": 49,
	"X": 50,
	"Y": 51,
	"Z": 52,
}

func calculatePriorityScore(value rune) (score int) {
	if unicode.IsUpper(value) {
		score = uppercaseMap[string(value)]
	} else {
		score = lowercaseMap[string(value)]
	}

	return score
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func determineDuplicateItems(bag Bag) (dupes []Item) {
	//secondItems := bag.secondCompartment
	secondItems := make([]Item, len(bag.secondCompartment))
	copy(secondItems, bag.secondCompartment)
	var alreadyAddedItems []string

	for _, item1 := range bag.firstCompartment {
		for index2, item2 := range secondItems {
			if item1.value == item2.value && !contains(alreadyAddedItems, item1.value) {
				dupes = append(dupes, item1)
				secondItems[index2] = Item{}
				alreadyAddedItems = append(alreadyAddedItems, item1.value)
				break
			}
		}
	}
	return dupes
}

func RemoveIndex(s []Item, index int) []Item {
	return append(s[:index], s[index+1:]...)
}

func splitBagDataToCompartments(fullBag string) (compartment1 string, compartment2 string) {
	firstCompartment := fullBag[0 : len(fullBag)/2]
	secondCompartment := strings.ReplaceAll(fullBag, firstCompartment, "")

	compartment1 = firstCompartment
	compartment2 = secondCompartment

	return compartment1, compartment2
}

func determineItemsInCompartmentString(compartmentString string) (items []Item) {
	chars := []rune(compartmentString)

	for _, rawItem := range chars {
		score := calculatePriorityScore(rawItem)

		processedItem := Item{
			value:         string(rawItem),
			priorityScore: score,
		}

		items = append(items, processedItem)
	}

	return items
}

func parseInputFile() (bags []Bag) {
	f, err := os.Open("./2022/Day3/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		bag := scanner.Text()

		firstString, secondString := splitBagDataToCompartments(bag)

		g := Bag{
			firstCompartment:  determineItemsInCompartmentString(firstString),
			secondCompartment: determineItemsInCompartmentString(secondString),
		}

		g.itemsInBothCompartments = determineDuplicateItems(g)

		bags = append(bags, g)
	}

	return bags
}

func DetermineGroupType(bags []Bag) (groupType string) {
	// append([]int{1,2}, []int{3,4}...)
	firstBag := append(bags[0].firstCompartment, bags[0].secondCompartment...)
	secondBag := append(bags[1].firstCompartment, bags[1].secondCompartment...)
	thirdBag := append(bags[2].firstCompartment, bags[2].secondCompartment...)

	for _, bag1 := range firstBag {
		for _, bag2 := range secondBag {
			if bag2.value == bag1.value {
				for _, bag3 := range thirdBag {
					if bag3.value == bag1.value {
						groupType = bag1.value
					}
				}
			}
		}
	}

	return groupType
}

func main() {
	sortedBags := parseInputFile()
	dupeScore := 0
	var groups []Group
	groupScore := 0

	for _, bag := range sortedBags {
		for _, score := range bag.itemsInBothCompartments {
			dupeScore += score.priorityScore
		}
	}

	// Sort bags into groups
	var bagsToadd []Bag
	count := 0
	for _, bag := range sortedBags {
		bagsToadd = append(bagsToadd, bag)
		count++
		if count == 3 {
			newGroup := Group{
				bags:      bagsToadd,
				groupType: DetermineGroupType(bagsToadd),
			}
			char := []rune(newGroup.groupType)
			newGroup.groupScore = calculatePriorityScore(char[0])
			groups = append(groups, newGroup)
			bagsToadd = nil
			count = 0
		}
	}

	for _, group := range groups {
		var duplicateItems []Item
		for _, bag := range group.bags {
			duplicateItems = append(duplicateItems, bag.itemsInBothCompartments...)
		}
		//groupItems := determineItemsInAllCompartmentsOfAllBags(duplicateItems, group.bags)
		groupItem := DetermineGroupType(group.bags)
		fmt.Println(groupItem)
		fmt.Println("Group score: ", group.groupScore)
		groupScore += group.groupScore
	}

	fmt.Println("Total dupe score: ", dupeScore)
	fmt.Println("Total group score: ", groupScore)
}
