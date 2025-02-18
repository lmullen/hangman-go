package main

import (
	"errors"

	tm "github.com/buger/goterm"
)

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
			if errors.Is(err, ErrBadGuess) {
				p.MakeGuess("", "Bad guess. Guess a single letter.")
				continue
			}
		}
		p.MakeGuess(guess, "")

		// Check if the puzzle is complete
		if p.Complete() {
			displayWon()
			tm.Println("The puzzle was:")
			tm.Println(tm.Color(p.Puzzle(), tm.GREEN))
			tm.Flush()
			break
		}

		if p.Lost() {
			displayLost()
			tm.Println("The answer was:")
			tm.Println(tm.Color(p.Answer(), tm.RED))
			tm.Println("\nYour final guess was:")
			tm.Println(tm.Color(p.Puzzle(), tm.GREEN))
			tm.Flush()
			break
		}
	}
}
