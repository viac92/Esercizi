// Calcola la somma dei primi n. quadrati 1^n, 2^n ecc

package main

import "fmt"

func main() {
	var n, s, i int
	fmt.Print("Inserire un numero conterò le somme dei quadrati fino al quadrato richiesto: ")
	fmt.Scan(&n)

	for i < n {
		i ++
		s += (i * i)
		fmt.Println("La", i, "somma è:", s)
	}
}
