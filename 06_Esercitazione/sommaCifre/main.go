// Stabilire se le somme delle cifre di un numero di tre cifre è maggiore di 10 

package main

import "fmt"

func main() {
	var num1 int
	fmt.Print("Inserisci un numero intero di 3 cifre: ")
	fmt.Scan(&num1)

	cUni := num1 % 10
	cDec := ((num1 % 100) - num1 % 10) / 10
	cCen := ((num1 % 1000) - num1 % 100) / 100
	somma := cUni + cDec + cCen
	

	fmt.Println(cUni)
	fmt.Println(cDec)
	fmt.Println(cCen)

	if somma > 10 {
		fmt.Println("La somma delle cifre di", num1, "che è uguale a", somma, "è maggiore di 10.")
	} else {
		fmt.Println("La somma delle cifre di", num1, "che è", somma, "è minore di 10.")
	}
}