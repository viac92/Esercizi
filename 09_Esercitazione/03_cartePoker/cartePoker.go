// Esercizio:
// Numerare carte da poker 13x4 = 52, numerarle da 0 a 51. Scrivere un programma che legge un numero da 0 a 51 e stampa la carta corrispondente. con simbolo ASCII.

package main

import (
	"fmt"
)

// Seme dato un numero restituisce il seme 
func Seme(n int) rune {
	switch n / 13 {
	case 0:
		return '\u2665'
	case 1:
		return '\u2666'
	case 2:
		return '\u2663'
	case 3:
		return '\u2660'
	default:
		return '?'
	}
}

// Carta dato un numero restituisce il valore della carta
func Carta(n int) string {
	switch n % 13 {
	case 0: 
		return "A"
	case 9:
		return "10"
	case 10: 
		return "J"
	case 11:
		return "Q"
	case 12:
		return "K"
	default:
		s := 49 + (n % 13)
		return string(s) 
	}
}

func main()	{
	for i := 0; i < 52; i++ {
	fmt.Println(Carta(i) + string(Seme(i)))
	}
}