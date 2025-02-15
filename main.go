package main

import "errors"

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
	}
}
