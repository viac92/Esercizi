// Esercizio 1
// Dato un numero intero stabilire se il numero contiene la cifra 3 (usare i cicli e il break).

package main

import (
	"fmt"
)

func main() {
	var n int
	var b bool

	fmt.Print("Inserisci un numero troverò se è presente la cifra 3: ")
	fmt.Scan(&n)

	p := n
	for n > 2 {
		if n % 10 == 3 {
			b = true
			break
		} else {
			n /= 10
		}
	}
	
	if b {
		fmt.Println("Il numero", p, "ha almeno un 3.")
	} else {
		fmt.Println("Nessun 3 in", p)
	}
}
