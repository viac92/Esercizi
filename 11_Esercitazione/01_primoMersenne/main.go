// Un primo di Mersenne è un numero primo della forma 2^x - 1. 
// scrivereun programma che stampi i primi n numeri di Mersenne. 

package main

import (
	"fmt"
	"math"
)

// dato un n restituisce l'n-esimo numero primo
func numPrimi(n int) int {
	var numero int = 2
	var index int
	var contaPrimi int	
	//ciclo per numeri primi
	for {
		//ciclo per successione numeri
		for {
			index = 2
			//ciclo per verifica numero primo
			for index < numero {
				if numero % index == 0 {
					break
				}	
				index++
			}
			if index >= numero {
				
				break
			}
			numero++
		} 

		contaPrimi++
		if contaPrimi == n {
			break
		}
		numero++
	}
	return numero
}

func numMersenne(n int) {
	m := math.Pow(2, float64(numPrimi(n))) - 1
	fmt.Println(m)

}

func main() {
	var n int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
	numMersenne(i + 1)
	}
}