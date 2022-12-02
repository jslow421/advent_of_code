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

func calculateOutcomeAndPoints(input string, opponent string) (outcome outcome, points int) {
	if input == opponent {
		outcome = tie
		points += 3
	}

	switch input {
	case "Y":
		// Played paper which beats rock
		points += 2
		if opponent == "A" {
			outcome = win
		}
	case "Z":
		// Played scissors which beats paper
		points += 3
		if opponent == "B" {
			outcome = win
		}
	case "X":
		// Played rock which beats scissors
		points += 1
		if opponent == "C" {
			outcome = win
		}
	}

	if outcome == win {
		points += 6
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
		/*		fmt.Println("Game outcome: ", game.outcome)
				fmt.Println("Points for game: ", game.points)
				fmt.Println("")*/
	}

	fmt.Println("Total wins: ", totalWins)
	fmt.Println("Total points: ", totalPoints)
}
