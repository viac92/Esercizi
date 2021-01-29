// Esercizio: 2
// Stampare le potenze di due minori o uguali a n.

package main

import (
	"fmt"
)

func main() {
	var n, p, p2 int

	fmt.Print("Inserisci un numero ti dirò le potenze di 2 minori o uguali: ")
	fmt.Scan(&n)
	p = 1
	p2 = 2

	if n != 1 {
		for p2 <= n {
			if p <= 1{
			fmt.Println("1")
			}
			p *= 2
			p2 *= 2 
			fmt.Println(p)
		}
	} else {
		fmt.Println("1")
	}	
}
