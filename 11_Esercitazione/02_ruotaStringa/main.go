// Scrivere una funzione che data una stringa e un intero ruoti la stringa di quel 
// numero di posizioni. es Tappeto, 3 -> petoTap.

package main

import (
	"fmt"
)


// dato n scambia le prime n lettere posizionandole alla fine della parola 
func ruotaStringa(n int, s string) (string) {
	var prefix, second string
	for i, char := range s {				
		if i < n {
			prefix += string(char)
		} else {
			second += string(char)
		}
	}
	return second + prefix
}

func main() {
	var n int
	var s string

	fmt.Print("Inserisci una parola: ")
	fmt.Scan(&s)
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)

	newS := ruotaStringa(n, s)

	fmt.Println(newS)
	
}
