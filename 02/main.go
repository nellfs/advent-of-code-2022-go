package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input, _ := os.Open("input")
	defer input.Close()
	inputScanner := bufio.NewScanner(input)

	var score int
	scores := map[string]int{"Rock Paper": 1, "C Y": 2, "A Z": 3, "A X": 4, "B Y": 5, "C Z": 6, "C X": 7, "A Y": 8, "B Z": 9}
	for inputScanner.Scan() {
		score += scores[inputScanner.Text()]
	}

	fmt.Println(score)
}
