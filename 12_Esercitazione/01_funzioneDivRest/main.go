// Scrivere una funzione che dati x e y restituisce 
// divisione intera x/y e resto come due valori restituiti.

package main

import "fmt"

func divisioneResto(num, div int) (quoz, rest int) {
	return num / div, num & div
}

func main()  {
	var n, d int
	fmt.Print("Inserisci il primo numero: ")
	fmt.Scan(&n)
	fmt.Print("Inserisci il secondo numero: ")
	fmt.Scan(&d)

	quoz, rest := divisioneResto(n, d)

	fmt.Print("Il quoziente tra ", n, " e ", d, " è ", quoz, ".\nIl resto è ", rest, ".\n")
}