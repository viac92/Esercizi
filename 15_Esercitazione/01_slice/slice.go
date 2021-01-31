// Esercizio: 1
// Scrivere un programma con una funzione che decide se un numero è primo.
// Nel main chiedere all'utente un valore n e inserire in una slice tutti i primi c = n.
// Stampare la slice, la sua lunghezza e la sua capacità.

package main

import (
	"fmt"
)

// primo restituisce n numeri primi
func primo(n int) []int {
	var numeriPrimi []int
	numeriPrimi = make([]int, 0, n)
	var flag bool = true 
	
	//ciclo numeri da testare
	for num := 2; num < n; num++ {
	
		//ciclo controllo
		for i := 2; i < num; i++ {
			if num % i == 0 {
				flag = false
				break
			}
		}

		if flag {
			numeriPrimi = append(numeriPrimi, num)
		}
		flag = true
	}

	return numeriPrimi
}

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Print(primo(n), cap(primo(n)), len(primo(n)), "\n")
}
