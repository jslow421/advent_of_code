package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Crate struct {
	label string
}

type Stack struct {
	index  int
	crates []Crate
}

// Global Stacks
var s1 = Stack{
	index: 0,
}
var s2 = Stack{
	index: 1,
}
var s3 = Stack{
	index: 2,
}
var s4 = Stack{
	index: 3,
}
var s5 = Stack{
	index: 4,
}
var s6 = Stack{
	index: 5,
}
var s7 = Stack{
	index: 6,
}
var s8 = Stack{
	index: 7,
}
var s9 = Stack{
	index: 8,
}

type Instruction struct {
	crateCount        int
	beginningLocation int
	endingLocation    int
}

var globalInstructions []Instruction

func RemoveIndex(s []Crate, index int) []Crate {
	return append(s[:index], s[index+1:]...)
}

func AddIndex(s []Crate, index int) []Crate {
	return append(s[:index+1], s[index:]...)
}

func moveCrates() {
	for index, todo := range globalInstructions {
		fmt.Println("Iteration: ", index)
		var globalStackMap = map[int]Stack{
			0: s1,
			1: s2,
			2: s3,
			3: s4,
			4: s5,
			5: s6,
			6: s7,
			7: s8,
			8: s9,
		}

		for i := 1; i <= todo.crateCount; i++ {
			var pullingFromStack = globalStackMap[todo.beginningLocation-1]
			var movingToStack = globalStackMap[todo.endingLocation-1]
			itemToMove := pullingFromStack.crates[0]
			if movingToStack.index == 0 {
				s1.crates = AddIndex(s1.crates, 0)
				s1.crates[0] = itemToMove
			}
			if movingToStack.index == 1 {
				s2.crates = AddIndex(s2.crates, 0)
				s2.crates[0] = itemToMove
			}
			if movingToStack.index == 2 {
				s3.crates = AddIndex(s3.crates, 0)
				s3.crates[0] = itemToMove
			}
			if movingToStack.index == 3 {
				s4.crates = AddIndex(s4.crates, 0)
				s4.crates[0] = itemToMove
			}
			if movingToStack.index == 4 {
				s5.crates = AddIndex(s5.crates, 0)
				s5.crates[0] = itemToMove
			}
			if movingToStack.index == 5 {
				s6.crates = AddIndex(s6.crates, 0)
				s6.crates[0] = itemToMove
			}
			if movingToStack.index == 6 {
				s7.crates = AddIndex(s7.crates, 0)
				s7.crates[0] = itemToMove
			}
			if movingToStack.index == 7 {
				s8.crates = AddIndex(s8.crates, 0)
				s8.crates[0] = itemToMove
			}
			if movingToStack.index == 8 {
				s9.crates = AddIndex(s9.crates, 0)
				s9.crates[0] = itemToMove
			}

			if pullingFromStack.index == 0 {
				s1.crates = RemoveIndex(s1.crates, 0)
			}
			if pullingFromStack.index == 1 {
				s2.crates = RemoveIndex(s2.crates, 0)
			}
			if pullingFromStack.index == 2 {
				s3.crates = RemoveIndex(s3.crates, 0)
			}
			if pullingFromStack.index == 3 {
				s4.crates = RemoveIndex(s4.crates, 0)
			}
			if pullingFromStack.index == 4 {
				s5.crates = RemoveIndex(s5.crates, 0)
			}
			if pullingFromStack.index == 5 {
				s6.crates = RemoveIndex(s6.crates, 0)
			}
			if pullingFromStack.index == 6 {
				s7.crates = RemoveIndex(s7.crates, 0)
			}
			if pullingFromStack.index == 7 {
				s8.crates = RemoveIndex(s8.crates, 0)
			}
			if pullingFromStack.index == 8 {
				s9.crates = RemoveIndex(s9.crates, 0)
			}
		}
	}
}

func populateStacks(stacksInLine string) {
	char := []rune(stacksInLine)
	var value []string
	places := []int{
		1, 5, 9, 13, 17, 21, 25, 29, 33,
	}

	for index, item := range char {
		for _, requiredValues := range places {
			if requiredValues == index {

				value = append(value, string(item))
			}
		}
	}

	for index, v := range value {
		if v == " " {
			continue
		}
		if index == 0 {
			s1.crates = append(s1.crates, Crate{label: v})
		}
		if index == 1 {
			s2.crates = append(s2.crates, Crate{label: v})
		}
		if index == 2 {
			s3.crates = append(s3.crates, Crate{label: v})
		}
		if index == 3 {
			s4.crates = append(s4.crates, Crate{label: v})
		}
		if index == 4 {
			s5.crates = append(s5.crates, Crate{label: v})
		}
		if index == 5 {
			s6.crates = append(s6.crates, Crate{label: v})
		}
		if index == 6 {
			s7.crates = append(s7.crates, Crate{label: v})
		}
		if index == 7 {
			s8.crates = append(s8.crates, Crate{label: v})
		}
		if index == 8 {
			s9.crates = append(s9.crates, Crate{label: v})
		}
	}
}

func getInstructions(input string) {
	if input != "" {
		instructions := strings.Fields(input)
		createCount, _ := strconv.Atoi(instructions[1])
		fromValue, _ := strconv.Atoi(instructions[3])
		toValue, _ := strconv.Atoi(instructions[5])

		newInstruction := Instruction{
			crateCount:        createCount,
			beginningLocation: fromValue,
			endingLocation:    toValue,
		}

		globalInstructions = append(globalInstructions, newInstruction)
	}
}

func parseInputFile() {
	f, err := os.Open("./2022/Day5/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		// These are the initial stacks and a lazy way to ignore the lines
		if scanner.Text() == "" || scanner.Text() == " 1   2   3   4   5   6   7   8   9" {
			break
		}
		populateStacks(scanner.Text())
	}

	for scanner.Scan() {
		getInstructions(scanner.Text())
	}
}

func main() {
	parseInputFile()

	fmt.Println("Completed stack 1: ", s1.crates)
	fmt.Println("Completed stack 2: ", s2.crates)
	fmt.Println("Completed stack 3: ", s3.crates)
	fmt.Println("Completed stack 4: ", s4.crates)
	fmt.Println("Completed stack 5: ", s5.crates)
	fmt.Println("Completed stack 6: ", s6.crates)
	fmt.Println("Completed stack 7: ", s7.crates)
	fmt.Println("Completed stack 8: ", s8.crates)
	fmt.Println("Completed stack 9: ", s9.crates)
	moveCrates()
	fmt.Println("")
	fmt.Println("Completed stack 1: ", s1.crates)
	fmt.Println("Completed stack 2: ", s2.crates)
	fmt.Println("Completed stack 3: ", s3.crates)
	fmt.Println("Completed stack 4: ", s4.crates)
	fmt.Println("Completed stack 5: ", s5.crates)
	fmt.Println("Completed stack 6: ", s6.crates)
	fmt.Println("Completed stack 7: ", s7.crates)
	fmt.Println("Completed stack 8: ", s8.crates)
	fmt.Println("Completed stack 9: ", s9.crates)
}
