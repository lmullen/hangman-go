package main

import (
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
