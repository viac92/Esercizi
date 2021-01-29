// Esercizio: 1
// Leggere una stringa e stampare quante "a" contiene.

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
	
	var contaA int
	for _, x := range s {
		if x == 97 {
			contaA++
		}
	}
	fmt.Print("In \"", s, "\" ci sono: ", contaA, " a.\n")

}
