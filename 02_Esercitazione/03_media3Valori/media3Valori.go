// Esercizio: 3
// Fare programma che calcola la media di tre valori letti in ingresso

package main

import (
	"fmt"
)

func main() {
	var num1, num2, num3 float64
	
	fmt.Print("Inserisci tre valori per fare la media: ")
	fmt.Scan(&num1, &num2, &num3)
	fmt.Println("La media è:", (num1 + num2 + num3) / 3)
}
