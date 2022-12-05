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
	bags []Bag
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

		//alreadyAddedItems = nil
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

func calculateTotalDupeScore(bag Bag) (score int) {
	totalScore := 0

	for _, dupeItem := range bag.itemsInBothCompartments {
		totalScore += dupeItem.priorityScore
	}

	return totalScore
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

func determineItemsInAllCompartmentsOfAllBags(duplicateItems []Item, bags []Bag) (items []Item) {
	hashedDupes := make(map[int]string)
	returnItems := make(map[int]Item)

	for _, item := range duplicateItems {
		hashedDupes[item.priorityScore] = item.value
	}

	for _, hashItem := range hashedDupes {
		for _, bag := range bags {
			for _, bagItem := range bag.itemsInBothCompartments {
				if hashItem == bagItem.value {
					//delete(hashedDupes, bagItem.priorityScore)
					//items = append(items, bagItem)
					returnItems[bagItem.priorityScore] = bagItem
					break
				}
			}
		}
	}

	for _, dedupedItems := range returnItems {
		items = append(items, dedupedItems)
	}

	return items
}

func main() {
	sortedBags := parseInputFile()
	dupeScore := 0
	var groups []Group
	//var duplicateItems []Item

	for _, bag := range sortedBags {
		/*		fmt.Println(index)
				fmt.Println(bag.firstCompartment)
				fmt.Println(bag.secondCompartment)
				fmt.Println(bag.itemsInBothCompartments)
				fmt.Println("")*/
		//duplicateItems = append(duplicateItems, bag.itemsInBothCompartments...)
		/*		fmt.Println("Dupe score: ", calculateTotalDupeScore(bag))
				fmt.Println("")*/
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
		fmt.Println("Count: ", count)
		if count == 3 {
			newGroup := Group{
				bags: bagsToadd,
			}
			groups = append(groups, newGroup)
			bagsToadd = nil
			count = 0
		}
	}

	for index, group := range groups {
		fmt.Println("Group index: ", index)
		fmt.Println(group)
		fmt.Println(group.bags)
		fmt.Println("")

	}

	//value := determineItemsInAllCompartmentsOfAllBags(duplicateItems, sortedBags)

	/*	for _, points := range value {
		fmt.Println(points)
		dupeScore += points.priorityScore
	}*/

	fmt.Println("Total dupe score: ", dupeScore)
}
