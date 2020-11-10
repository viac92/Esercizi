// Scrivere una funzione uguale alla 1 ma i due valori vengono 
// scritti in due variabili passate come puntatori.

package main

import "fmt"

var quoz, rest *int

func divRest(q, d int) {
	*quoz = q / d
	*rest = q % d	
}

func main()  {
	var q, d int

	fmt.Print("Inserisci il primo numero: ")
	fmt.Scan(&q)
	fmt.Print("Inserisci il secondo numero: ")
	fmt.Scan(&d)
	
	quoz = &q
	rest = &d

	divRest(q, d)
	fmt.Println(*quoz, *rest)
}