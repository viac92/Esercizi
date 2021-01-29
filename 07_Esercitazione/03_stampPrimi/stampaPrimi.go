// Esercizio:
// Dato n stapare i primi n numeri primi.

package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)

	var numPrimo int = 2

	for i := 0; i < n; i++ {

		// stapa il numero primo
		cercaPrimo := 2

		for {
			for {
				if numPrimo%cercaPrimo == 0 {
					break
				}

				if cercaPrimo == numPrimo {
					break
				}
				cercaPrimo++
			}

			if cercaPrimo != numPrimo {
				numPrimo++
			} else {
				fmt.Println(numPrimo)
				numPrimo++
				break
			}
		}
	}
}
