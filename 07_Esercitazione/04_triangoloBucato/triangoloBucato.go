// Esercizio: 
// Disegnare il triangolo bucato.
// *****
// *  *
// * *
// *


package main

import (
	"fmt"
)

func main() {
	var n int 
	
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)
	
	// ciclo per le righe
	for i := 0; i < n; i++ { 		
		if i == n - 1 {
			fmt.Println("*")
			break
		}

		if i == 0 {
			for in := 0; in < n; in++ {
				fmt.Print("*")
			}
		} else {
				for {
					fmt.Print("*")
						for x := 0; x < (n - i) - 2; x++ {
							fmt.Print(" ")
						}
					fmt.Print("*")
					break	
				}
			}
		fmt.Println()	
	}
}