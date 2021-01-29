// Esercizio: 3
// Data una stringa ASCII, stampare tutto ciò che compare prima della prima '/'.

package main

import (
		"fmt"
		"bufio"
		"os"
)

func main() {
	fmt.Print("Inserisci una stringa: ")
	
	b := bufio.NewScanner(os.Stdin)
	
	b.Scan()
	s := b.Text()

	for _, x := range s {
		fmt.Print(string(x))
		if x == 47 {
			break
		}
	}
	fmt.Println()
}
