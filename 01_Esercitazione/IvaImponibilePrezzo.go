// Esercizio: 1
// Data una certa data calcolare quanti giorni mancano al primo gennaio e quanti giorni mancano a natale facendo finta che tutti i mesi siano di 30 giorni.

package main

import (
	"fmt"
)

func main() {
	var g, m, mNatale, mGennaio int
	
	fmt.Print("Inserisci la data odierna nel seguente formato g/m - ")
	fmt.Scan(&g)
	fmt.Scan(&m)
	
	mNatale = (30 - g) + 30*(12-1-m) + 25
	mGennaio = mNatale + 6
	
	fmt.Println("Mancano", mNatale, "giorni a Natale.")
	fmt.Println("Mancano", mGennaio, "giorni al primo di Gennaio.")
}
