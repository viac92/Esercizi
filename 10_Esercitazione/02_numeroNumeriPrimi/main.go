// 2 Scrivere un programma che dato n calcola il numero di numeri primi minori o uguali a n. Usarlo per verificare sperimentalmente che questo P(n) è asintoti

package main

import "fmt"

func main() {
	var n int 
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)
	
	if n < 2 {
		fmt.Println("Valore non valido.")
		return
	}

	var flag bool
	var numeroPrimi int

	for i := 0; i < n; i++ {
		var k int
		k = 2
		if n % k == 0 {
			break
		}
	}

	if flag {
		numeroPrimi++
	}
	fmt.Println(numeroPrimi)
}