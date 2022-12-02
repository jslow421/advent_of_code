package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Game struct {
	opponentPlay string
	yourPlay     string
	outcome      outcome
	points       int
}

type outcome int

const (
	lose outcome = iota
	tie
	win
)

var equivalentMap = map[string]string{
	"A": "X",
	"B": "Y",
	"C": "Z",
}

var winMap = map[string]string{
	"A": "Y",
	"B": "Z",
	"C": "X",
}

var drawMap = map[string]string{
	"A": "X",
	"B": "Y",
	"C": "Z",
}

var loseMap = map[string]string{
	"A": "Z",
	"B": "X",
	"C": "Y",
}

func determineCorrectChoice(opponentChoice string, yourChoiceKey string) (translatedChoice string) {

	switch yourChoiceKey {
	case "X":
		// You should lose
		translatedChoice = loseMap[opponentChoice]
	case "Y":
		// You should draw
		translatedChoice = drawMap[opponentChoice]
	case "Z":
		// You should win
		translatedChoice = winMap[opponentChoice]
	}

	return translatedChoice
}

func calculateOutcomeAndPoints(input string, opponent string) (outcome outcome, points int) {

	translatedChoice := determineCorrectChoice(opponent, input)

	switch translatedChoice {
	case "Y":
		// Played paper which beats rock
		points += 2
		if opponent == "A" {
			outcome = win
		} else if opponent == "B" {
			outcome = tie
		}
	case "Z":
		// Played scissors which beats paper
		points += 3
		if opponent == "B" {
			outcome = win
		} else if opponent == "C" {
			outcome = tie
		}
	case "X":
		// Played rock which beats scissors
		points += 1
		if opponent == "C" {
			outcome = win
		} else if opponent == "A" {
			outcome = tie
		}
	}

	if input == opponent {
		outcome = tie
		points += 3
	}

	switch outcome {
	case win:
		points += 6
	case tie:
		points += 3
	}

	return outcome, points
}

func parseInputFile() (games []Game) {
	f, err := os.Open("./2022/Day2/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		plays := strings.Fields(scanner.Text())
		opponentPlay := plays[0]
		yourPlay := plays[1]
		outcome, points := calculateOutcomeAndPoints(yourPlay, opponentPlay)

		g := Game{
			opponentPlay: opponentPlay,
			yourPlay:     yourPlay,
			outcome:      outcome,
			points:       points,
		}

		games = append(games, g)
	}

	return games
}

func main() {
	games := parseInputFile()
	var totalWins int
	var totalPoints int

	for _, game := range games {
		if game.outcome == win {
			totalWins++
		}
		totalPoints += game.points
	}

	fmt.Println("Total games: ", len(games))
	fmt.Println("Total wins: ", totalWins)
	fmt.Println("Total points: ", totalPoints)
}
