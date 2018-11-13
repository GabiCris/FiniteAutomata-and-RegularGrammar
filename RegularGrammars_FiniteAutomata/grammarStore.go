package main

import (
	"unicode"
)

type RegGrammar struct {
	terminalS    map[rune]bool
	nonTerminalS map[rune]bool
	productions  map[rune]ProductionResult
}

type ProductionResult [][]rune

var terminalSet = make(map[rune]bool)
var nonTerminalSet = make(map[rune]bool)
var productionsMap = make(map[rune]ProductionResult)

var obtainedGrammar = RegGrammar{terminalSet, nonTerminalSet, productionsMap}

func addProduction(pMap map[rune]ProductionResult, lh rune, rh []rune) map[rune]ProductionResult {
	pMap[lh] = append(pMap[lh], rh)
	saveTerminals(lh, rh)
	return pMap
}

func saveTerminals(lh rune, rh []rune) {
	if unicode.IsUpper(lh) {
		nonTerminalSet[lh] = true
	} else {
		panic("Invalid Grammar")
		terminalSet[lh] = true
	}

	for _, r := range rh {
		if unicode.IsUpper(r) {
			nonTerminalSet[r] = true
		} else {
			terminalSet[r] = true
		}
	}
}

func checkRegGrammar() bool {
	isRegular := true
	for key, v := range productionsMap {
		for _, val := range v {
			if len(val) > 2 {
				isRegular = false
				panic("Grammar has more than 2 symbols in the rhs!")
				break
			} else {
				if len(val) == 1 {
					if !unicode.IsLower(val[0]) {
						isRegular = false
						panic("Grammar should have either a single terminal or a terminal followed by a non-terminal on the rhs!")
						break
					}
					if val[0] == 'ε' {
						if key != 'S' {
							isRegular = false
							panic("Transitions to epsilon are only accepted from the initial state S!")
							break
						} else if ensureS() == false {
							isRegular = false
							panic("If S->ε, S cannot appear in any production on the rhs!")
							break
						}
					}
				}
				if len(val) == 2 && (!unicode.IsLower(val[0]) || !unicode.IsUpper(val[1])) {
					isRegular = false
					panic("Grammar should have either a single terminal or a terminal followed by a non-terminal on the rhs!")
					break
				}
			}
		}
	}
	return isRegular
}

func ensureS() bool {
	for _, v := range productionsMap {
		for _, list := range v {
			for _, el := range list {
				if el == 'S' {
					return false
				}
			}
		}
	}
	return true
}
