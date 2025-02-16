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
	tm.MoveCursor(1, 1)
	tm.Println("Welcome to ...")
	tm.Print(titleText)
	tm.Println()
	tm.Print("by Paul and Daddy\n\n\n")
	tm.Println()
	tm.Println("Press return/enter to play ...")
	tm.Flush()
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}

//go:embed instructions.txt
var instructions string

func getInput() (string, error) {
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
	tm.Println(tm.Color(fmt.Sprintf("You have %v mistakes left.", p.mistakesLeft), tm.RED))
	tm.Println("You've guessed: ")
	tm.Println(p.Guesses())
	tm.Println()
	tm.Print("Here's the puzzle:\n\n")
	tm.Println(tm.Color(p.Puzzle(), tm.GREEN))
	tm.Println()
	if p.message != "" {
		tm.Println(tm.Color(p.message, tm.CYAN))
	} else {
		tm.Println()
	}
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

//go:embed won.txt
var won string

func displayWon() {
	interval := 200 * time.Millisecond
	count := 10
	for i := 0; i < count; i++ {
		tm.Clear()
		tm.MoveCursor(1, 1)
		tm.Print(won)
		tm.Printf("\n\n\n\n")
		tm.Flush()
		time.Sleep(interval)
		tm.Clear()
		tm.Flush()
		time.Sleep(interval)
	}
	tm.MoveCursor(1, 1)
	tm.Println(won)
	tm.Printf("\n\n\n\n")
	tm.Flush()
}

//go:embed lost.txt
var lost string

func displayLost() {
	tm.Clear()
	tm.MoveCursor(1, 1)
	tm.Print(lost)
	tm.Printf("\n\n\n\n")
	tm.Flush()
	time.Sleep(3 * time.Second)
}
