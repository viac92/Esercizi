// Esercizio:
// Fare programma che calcola la media di due valori letti in ingresso

package main

import "fmt"

func main() {
	var num1, num2 float32
	
	fmt.Print("Inserisci due numeri per fare la media: ")
	fmt.Scan(&num1, &num2)
	fmt.Println("La media è", (num1 + num2) / 2)
}
