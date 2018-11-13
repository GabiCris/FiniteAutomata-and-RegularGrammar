package main

import (
	"fmt"
)

func processGrammar() {
	fmt.Print("Choose input type:\n1. File\n2. Console\nYour choice: ")
	var line string
	fmt.Scanln(&line)
	if line == "1" {
		readFile()
	} else if line == "2" {
		readKeyboard()
	} else {
		fmt.Println("Invalid")
	}

	isReg := checkRegGrammar()
	fmt.Println("")
	if isReg {
		fmt.Println("Given grammar is Regular!")
	} else {
		fmt.Println("Given grammer is NOT Regular!")
	}
}

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("")
			fmt.Printf("Panic: %+v\n", r)
		}
	}()

	//processGrammar()
	//	readKeyboardFA()

	processGrammar()
	fmt.Println("")
	fmt.Println("Grammar converted to FA:")
	convertGrFa(obtainedGrammar)
	printAut(outputFA)
	fmt.Println("")

	fmt.Println("FA read from File:")
	readFileFA()
	convertFaGr(finiteAut)
	fmt.Println("")
	fmt.Println("FA converted to GRAMMAR:")
	printOutGrammar(outputGrammar)
}
