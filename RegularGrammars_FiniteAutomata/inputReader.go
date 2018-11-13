package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readFile() {
	file, err := os.Open("sample_grammar")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "->")
		rh := []rune(line[1])

		addProduction(productionsMap, rune(line[0][0]), rh)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	printGrammar()
}

func printGrammar() {
	fmt.Println("Productions: ")
	for key, val := range productionsMap {
		fmt.Printf("%c: %c\n", key, val)
	}

	fmt.Println("Non-terminals:")
	for key, _ := range nonTerminalSet {
		fmt.Printf("%c ", key)
	}
	fmt.Println("")

	fmt.Println("Terminals:")
	for key, _ := range terminalSet {
		fmt.Printf("%c ", key)
	}

}

func printOutGrammar(gr RegGrammar) {
	fmt.Println("Productions: ")
	for key, val := range gr.productions {
		fmt.Printf("%c: %c\n", key, val)
	}

	fmt.Println("Non-terminals:")
	for key, _ := range gr.nonTerminalS {
		fmt.Printf("%c ", key)
	}
	fmt.Println("")

	fmt.Println("Terminals:")
	for key, _ := range gr.terminalS {
		fmt.Printf("%c ", key)
	}
}

func readKeyboard() {
	fmt.Println("Input Productions:")
	var lineR string
	var linePrev string
	for {
		fmt.Scanln(&lineR)
		line := strings.Split(lineR, "->")
		rh := []rune(line[1])

		addProduction(productionsMap, rune(line[0][0]), rh)
		if lineR == linePrev {
			break
		}
		linePrev = lineR
	}
	printGrammar()

}

func readFileFA() {
	file, err := os.Open("fa-sample")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	finiteAut.states = strings.Split(scanner.Text(), " ")

	scanner.Scan()
	finiteAut.alphabet = strings.Split(scanner.Text(), " ")

	scanner.Scan()
	finiteAut.initState = scanner.Text()

	scanner.Scan()
	finiteAut.finalStates = strings.Split(scanner.Text(), " ")

	scanner.Scan()
	finiteAut.transitions = getTransitions(scanner.Text())

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	printAut(finiteAut)
}

func printAut(finiteAut FA) {
	fmt.Println("States: ")
	fmt.Println(finiteAut.states)

	fmt.Println("Alphabet: ")
	fmt.Println(finiteAut.alphabet)

	fmt.Println("Initial state: " + finiteAut.initState)
	fmt.Println("Final states: ")
	fmt.Println(finiteAut.finalStates)

	fmt.Println("Transitions:")
	for _, tr := range finiteAut.transitions {
		fmt.Print(tr.from + "->")
		for _, sym := range tr.symbols {
			fmt.Printf("%c", sym)
		}
		fmt.Print("->" + tr.to + "\n")
	}
}

func readKeyboardFA() {
	var lineR string

	fmt.Print("Input states: ")
	fmt.Scanln(&lineR)
	fmt.Println(lineR)
	finiteAut.states = strings.Split(lineR, " ")

	fmt.Print("Input Alphabet: ")
	fmt.Scanln(&lineR)
	finiteAut.alphabet = strings.Split(lineR, " ")

	fmt.Println("")
	fmt.Print("Input initial state: ")
	fmt.Scanln(&lineR)
	finiteAut.initState = lineR

	fmt.Println("")
	fmt.Print("Input final states: ")
	fmt.Scanln(&lineR)
	finiteAut.finalStates = strings.Split(lineR, " ")

	fmt.Println("")
	fmt.Print("Transitions: ")
	fmt.Scanln(&lineR)
	finiteAut.transitions = getTransitions(lineR)

	printAut(finiteAut)

}
