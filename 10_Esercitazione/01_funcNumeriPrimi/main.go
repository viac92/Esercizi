// numerare carte da poker 13x4 = 52, numerarle da 0 a 51. Scrivere un programma che legge un numero da 0 a 51 e stampa la carta corrispondente. con simbolo ASCII

package main

import "fmt"

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