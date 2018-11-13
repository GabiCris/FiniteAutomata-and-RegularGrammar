package main

var outputFA FA

var outputGrammar RegGrammar

func convertGrFa(grammar RegGrammar) FA {
	// keys := make([]string, len(grammar.nonTerminalS)+len(grammar.terminalS))
	// alphabet := make([]string, len(grammar.nonTerminalS))
	var keys []string
	var alphabet []string
	for key, _ := range grammar.nonTerminalS {
		keys = append(keys, string(key))
		alphabet = append(alphabet, string(key))
	}
	for key, _ := range grammar.terminalS {
		alphabet = append(alphabet, string(key))
	}
	alphabet = append(alphabet, "F")

	outputFA.states = append(keys, "F")

	var finalS []string
	finalS = append(finalS, "F")

	for key, val := range grammar.productions {
		for _, prod := range val {
			if prod[0] == 'ε' {
				finalS = append(finalS, string(key))
			}
		}
	}
	outputFA.finalStates = finalS

	var transitions []Transition

	for key, val := range grammar.productions {
		for _, prod := range val {
			if len(prod) == 2 {
				sym := []rune{prod[0]}
				transitions = append(transitions, Transition{string(key), string(prod[1]), sym})
			} else if len(prod) == 1 && prod[0] != 'ε' {
				sym := []rune{prod[0]}
				transitions = append(transitions, Transition{string(key), "F", sym})
			}
		}
	}

	outputFA.alphabet = alphabet
	outputFA.transitions = transitions
	outputFA.initState = "S"

	return outputFA
}

func convertFaGr(fa FA) RegGrammar {
	var nonTerminalS = make(map[rune]bool)

	for _, state := range fa.states {
		nonT := []rune(state)
		nonTerminalS[nonT[0]] = true
	}

	outputGrammar.nonTerminalS = nonTerminalS

	var terminals = make(map[rune]bool)
	for _, sym := range fa.alphabet {
		runeSym := []rune(sym)
		//	if unicode.IsLower(runeSym[0]) {
		terminals[runeSym[0]] = true
		//}
	}
	outputGrammar.terminalS = terminals

	var productions = make(map[rune]ProductionResult)
	for _, val := range fa.transitions {
		if contains(fa.finalStates, val.to) {
			from := []rune(val.from)
			to := []rune(val.to)
			sym := val.symbols
			sym = append(sym, to[0])
			productions[from[0]] = append(productions[from[0]], val.symbols)
			productions[from[0]] = append(productions[from[0]], sym)
		} else {
			from := []rune(val.from)
			to := []rune(val.to)
			sym := val.symbols
			sym = append(sym, to[0])
			productions[from[0]] = append(productions[from[0]], sym)
		}
	}

	outputGrammar.productions = productions
	return outputGrammar

}
