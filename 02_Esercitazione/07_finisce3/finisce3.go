// Esercizio:
// Dato un numero capire se finisce per tre.

package main

import (
	"fmt"
)

func main() {
	var num1 int
	
	fmt.Print("Inserire un numero: ")
	fmt.Scan(&num1)

	if num1%10 == 3 {
		fmt.Println("Il numero", num1, "finisce per 3.")
		} else {
			fmt.Println("Il numero", num1, "non finisce per 3.")
	}
}
