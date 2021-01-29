// Esercizio:
// Scrivere una funzione isPrime che dice se un numero è primo.

package main

import (
	"fmt"
)

// isPrime restituisce true se un numero è primo, false altrimenti
func isPrime(n int) bool {
	if n == 0 || n == 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n % i == 0 || n == 1 {
			return false	
		}
	}
	return true 
}

func main()  {
	var n int
	
	fmt.Print("Inserisci un numero ti dico se è primo: ")
	fmt.Scan(&n)
	fmt.Println(n, "è primo? -->", isPrime(n))
}
