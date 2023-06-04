package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	input, _ := os.Open("input")
	inputScanner := bufio.NewScanner(input)

	maxCalories := 0
	currentElfCalories := 0

	for inputScanner.Scan() {
		snack, err := strconv.Atoi(inputScanner.Text())
		currentElfCalories += snack

		if err != nil {
			if currentElfCalories > maxCalories {
				maxCalories = currentElfCalories
			}

			currentElfCalories = 0
		}
	}
	fmt.Println(maxCalories)

	input.Close()
}
