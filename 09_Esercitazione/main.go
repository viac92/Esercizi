// dato n stampare un triangolo nell'altra direzione 
// ****
//  ***
//   **
//    *

package main

import "fmt"

func main() {
	var n int
	fmt.Print("Inserisci il numero di * per la base di un triangolo decrescente: ")
	fmt.Scan(&n)
	
	for i := 0; i < n; i++ {
		for y := 0; y < i; y++ {
			fmt.Print(" ")
		}	
		for x := 0; x < n - i; x++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
