package main

import (
	"errors"
	"strings"
	"unicode"
)

type letter struct {
	letter string
	show   bool
}

type puzzle []*letter

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
		p = append(p, l)
	}
	return &p
}

func (p *puzzle) String() string {
	var sb strings.Builder
	for _, l := range *p {
		sb.WriteString(l.String())
	}
	return sb.String()
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
