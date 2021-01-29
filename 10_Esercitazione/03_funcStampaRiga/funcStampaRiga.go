// Esercizio:
// Come programma 2 con una singola funzione stampaRiga (che prende due parametri).

package main

import (
		"fmt"
		"io"
)

// asterischi dato n in input restituisce n *
func asterischi(n int) string {
	var out string
		for i := 0; i < n; i++ {
			out += "*"
		}
	return out
}

// spazi dato n restituisce spazi 0, 1
func spazi(n int) string {
	var out string
	for i := 0; i < n; i++ {
		out += " "
		} 
	return out
}

// stampaRiga dati due input s e a stampa s spazi e a asterischi 
func stampaRiga(s int, a int) {
	fmt.Print(spazi(s) + asterischi(a) + "\n")
}

func main()  {
var s, a int

	for	{
		fmt.Println("Inserisci due numeri, il primo per il numero di spazi, il secondo per gli asterischi: ")
		
		_, err := fmt.Scan(&s)
		if err == io.EOF {
			break
		}
		fmt.Scan(&a)
		stampaRiga(s, a)
	}	
}
