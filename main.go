package main

import "fmt"

func main() {
	displayTitle()
	input, err := getInput()
	p := NewPuzzle(input)
	if err != nil {
		panic(err)
	}
	for {
		input, err = displayPuzzle(p)
		if err != nil {
			panic(err)
		}
		guess, err := validateGuess(input)
		if err != nil {
			panic(err)
		}
		fmt.Println(guess)
	}
}
