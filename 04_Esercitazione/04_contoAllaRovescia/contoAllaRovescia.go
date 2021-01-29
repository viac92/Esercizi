// Esercizio:
// Stampa un conto alla rovescia partendo da un numero n.

package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Print("Inserisci il numero di partenza per il conto alla rovescia: ")
	fmt.Scan(&n)

	for i := 0; i < n; {
		fmt.Println(n)
		n--
		if n < 1 {
			fmt.Println("Buon anno!")
		}
	}
}
