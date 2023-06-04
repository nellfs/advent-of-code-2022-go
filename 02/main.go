package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.Open("input")
	defer input.Close()
	inputScanner := bufio.NewScanner(input)

	options := map[string]string{
		"A": "rock",
		"B": "paper",
		"C": "scissors",
		"X": "rock",
		"Y": "paper",
		"Z": "scissors",
	}

	var score int
	for inputScanner.Scan() {
		match := strings.Split(inputScanner.Text(), " ")
		if match[0] != "" {
			opponent, own := options[match[0]], options[match[1]]
			score += PlayRockPaperScissors(opponent, own)
		}
	}
	fmt.Println(score)
}

func PlayRockPaperScissors(opponentChoice, ownChoice string) int {
	points := map[string]int{
		"rock":     1,
		"paper":    2,
		"scissors": 3,
	}

	if ownChoice == opponentChoice {
		return 3 + points[ownChoice]
	} else if (points[ownChoice]-points[opponentChoice]+3)%3 == 1 {
		return 6 + points[ownChoice]
	} else {
		return 0 + points[ownChoice]
	}
}
