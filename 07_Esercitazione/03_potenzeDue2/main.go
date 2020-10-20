// Stampa le prime n potenze di 2

package main

import "fmt"

func main() {
	var n, p int
	fmt.Print("Inserisci il numero di potenze di due che vuoi stampare: ")
	fmt.Scan(&n)

	p = 2

	for i := 0; i <= n - 1; {
		if i == 0 {
			p = 1
		}
		p *= 2	
		fmt.Println(i + 1, "-", p)
		i++
	} 
}