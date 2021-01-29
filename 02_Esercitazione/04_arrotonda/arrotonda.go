// Esercizio: 4
// Dato un numero stamparlo arrotondato all'intero.

package main

import ("fmt"
		"math"
)

func main() {
	var num1, intNum, dec float64
	
	fmt.Print("Inserire un numero con la virgola: ")
	fmt.Scan(&num1)
	intNum = math.Floor(num1)
	dec = num1 - intNum

	if dec < 0.4 {
		fmt.Println("Il numero arrotondato è", math.Floor(num1))
	} else {
		fmt.Println("Il numero arrotondato è:", math.Ceil(num1))
	}
}
