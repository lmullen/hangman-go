package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	_ "embed"

	tm "github.com/buger/goterm"
)

//go:embed title.txt
var titleText string

// Display the title
func displayTitle() {
	tm.Clear()
	tm.MoveCursor(1, 2)
	tm.Println("  Welcome to ...")
	tm.Print(titleText)
	tm.Println()
	tm.Print("  by Paul and Daddy\n\n\n")
	tm.Println()
	tm.Println("  Press return/enter to play ...")
	tm.Print("\n\n\n  ")
	tm.Flush()
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}

func getInput() (string, error) {
	instructions := `
	Please type the hangman puzzle you want your friend to solve.
	Make sure they aren't looking! 
	Press return/enter once you have finished typing your puzzle.
	`
	tm.Clear()
	tm.MoveCursor(1, 1)
	tm.Println(instructions)
	tm.Flush()
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	input = strings.Replace(input, "\n", "", -1) //remove new line
	input = strings.TrimSpace(input)
	fmt.Println("You entered:")
	fmt.Println(input)
	return input, nil
}

func displayPuzzle(p *puzzle) (string, error) {
	tm.Clear()
	tm.MoveCursor(1, 1)
	tm.Println(tm.Color(p.message, tm.RED))
	tm.Println()
	tm.Println("You've guessed: ")
	tm.Println(p.Guesses())
	tm.Println()
	tm.Print("Here's the puzzle:\n\n")
	tm.Println(tm.Color(p.Puzzle(), tm.GREEN))
	tm.Println()
	tm.Println("Guess a single letter and press return/enter:")
	tm.Flush()
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	input = strings.TrimSpace(input)
	return input, nil
}
