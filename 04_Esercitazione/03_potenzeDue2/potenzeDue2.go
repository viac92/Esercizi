// Esercizio:
// Stampa le prime n potenze di 2.

package main

import (
	"fmt"
)

func main() {
	var n, p int
	
	fmt.Print("Inserisci il numero di potenze di due che vuoi stampare: ")
	fmt.Scan(&n)

	p = 1

	for p <= n {
		fmt.Println(p)
		p *= 2	
	} 
}
