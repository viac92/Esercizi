// Scrivere un programma calc.go che chiamato da console restituisca le seguenti operazioni
// ./calc 12 + 5
// 17
// ./calc 13 / 6
// 2.1
// operatori +, -, x, /

package main

import (
	"fmt"
	"os"
	"strconv"
)

func calc(op string, n1 int, n2 int) (int, bool) {
	switch op {
	case "+":
		return n1 + n2, true
	case "-":
		return n1 - n2, true
	case "x":
		return n1 * n2, true
	case "/":
		return n1 / n2, true
	}
	return 0, false
}

func main() {
	n1, err1 := strconv.Atoi(os.Args[1])
	n2, err2 := strconv.Atoi(os.Args[3])
	
	operazione, ok := calc(os.Args[2], n1, n2) 

	if ok && err1 == nil && err2 == nil {
	fmt.Printf("Il risultato è %d\n", operazione)
	} else {
		fmt.Println("Errore")
	}
}
