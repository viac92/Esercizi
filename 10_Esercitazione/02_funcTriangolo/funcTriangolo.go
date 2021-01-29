// Esercizio: 2
// Scrivere un programma per stampare un tirangolo di asterischi.
// usando due funzioni per spazi ed asterischi.
//
// Es
// ******
//  *****
//   ****
//    ***
//     **
//      * 

package main

import (
	"fmt"
)

// asterischi dato n in input restituisce n *
func asterischi(n int, num int) string {
	var out string
		for i := 0; i < num - n; i++ {
			out += "*"
		}
	return out
}

// spazi dato n restituisce spazi 0, 1
func spazi(n int) string {
	var out string
	for i := 0; i < n; i++ {
		out += " "
	} 
	return out
}


func main() {
	var n int
	
	fmt.Print("Inserisci un numero: ")

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
	fmt.Print(spazi(i) + asterischi(i, n) + "\n")
	}

}
