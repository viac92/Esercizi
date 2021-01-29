// Esercizio: 2
// Scrivere una funzione uguale alla 1 ma i due valori vengono scritti in due variabili passate come puntatori.

package main

import (
	"fmt"
)


// divRest restituisce tramite puntatori in quo e rest il quoziente e il resto di q/d
func divRest(q, d *int) {
	quo := *q / *d
	rest := *q % *d	
	*q = quo
	*d = rest
}

func main()  {
	var q, d int

	fmt.Print("Inserisci il primo numero: ")
	fmt.Scan(&q)
	fmt.Print("Inserisci il secondo numero: ")
	fmt.Scan(&d)
	

	divRest(&q, &d)
	fmt.Println(q, d)
}
