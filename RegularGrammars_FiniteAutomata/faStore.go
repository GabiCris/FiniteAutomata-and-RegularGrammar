package main

import (
	"strings"
)

type FA struct {
	states      []string
	alphabet    []string
	initState   string
	finalStates []string
	transitions []Transition
}

type Transition struct {
	from    string
	to      string
	symbols []rune
}

var finiteAut FA

func getTransitions(line string) []Transition {
	lineSep := strings.Split(line, " ")
	lineSep = lineSep[:len(lineSep)-1]
	var transitions []Transition

	for _, el := range lineSep {
		symbols := strings.Split(el, ",")
		if len(symbols) > 2 {
			transitions = append(transitions, Transition{symbols[0], symbols[2], []rune(symbols[1])})
		} else {
			panic("Transition with no symbol!")
		}
	}
	return transitions
}
