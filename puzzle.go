package main

import (
	"errors"
	"fmt"
	"slices"
	"strings"
	"unicode"
)

type letter struct {
	letter string
	show   bool
}

type puzzle struct {
	letters      []*letter
	guesses      []string
	message      string
	mistakesLeft int
}

func NewLetter(r rune) *letter {
	show := false
	// Show non
	if !unicode.IsLetter(r) {
		show = true
	}
	return &letter{letter: string(r), show: show}
}

func (l *letter) String() string {
	if l.show {
		return l.letter
	}
	return "_"
}

func NewPuzzle(input string) *puzzle {
	var p puzzle
	var l *letter
	runes := []rune(input)
	for _, r := range runes {
		l = NewLetter(r)
		p.letters = append(p.letters, l)
	}
	p.mistakesLeft = 5
	return &p
}

func (p *puzzle) Puzzle() string {
	var sb strings.Builder
	for _, l := range p.letters {
		sb.WriteString(l.String())
	}
	return sb.String()
}

func (p *puzzle) Guesses() string {
	var sb strings.Builder
	for _, g := range p.guesses {
		sb.WriteString(g)
		sb.WriteString(", ")
	}
	return sb.String()
}

func (p *puzzle) MakeGuess(g string, msg string) {
	good := false
	p.message = msg
	// If we get a bad guess we don't need to check for anything
	if g == "" {
		return
	}
	g = strings.ToLower(g)
	if slices.Contains(p.guesses, g) {
		return
	}
	for _, l := range p.letters {
		if strings.EqualFold(l.letter, g) {
			l.show = true
			good = true
		}
	}
	p.guesses = append(p.guesses, g)
	if !good {
		p.mistakesLeft--

		p.message = fmt.Sprintf("You guessed wrong!")

	}
}

// ErrBadGuess means we have not received a single rune as a guess
var ErrBadGuess = errors.New("bad guess")

func validateGuess(s string) (string, error) {
	runes := []rune(s)
	if len(runes) != 1 || !unicode.IsLetter(runes[0]) {
		return "", ErrBadGuess
	}
	return string(runes[0]), nil
}
