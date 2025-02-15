package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

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
	tm.Flush()
	time.Sleep(5 * time.Second)
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
	tm.Print("Here's the puzzle:\n\n")
	tm.Println(p.String())
	tm.Println()
	tm.Println("Guess a single letter and press return/enter:")
	tm.Flush()
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	// input = strings.Replace(input, "\n", "", -1) //remove new line
	input = strings.TrimSpace(input)
	return input, nil
}
