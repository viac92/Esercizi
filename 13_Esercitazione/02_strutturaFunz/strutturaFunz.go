// Esercizio: 2
// Scrivere una funzione che prenda tre: argomenti un puntatore a una data, una stringa s e un intero x.
// La funzione deve modificare uno dei tre campi della data (il giorno se s = g, il mese se s = "m") ponendolo uguale a x.

package main

import (
	"fmt"
	"strconv"
)

// data memorizza una data dell'anno
type data struct {
	g, m, a int
}

// setData modifica il valore di un parametro di una data puntata
func setData(d *data, s string, x int) {
	switch s {
	case "g":
		d.g = x	
	case "m":
		d.m = x
	case "a":
		d.a = x
	}
}

// stringa stampa un tipo data in formato string
func stringa(d data) string {
	return strconv.Itoa(d.g) + "/" + strconv.Itoa(d.m) + "/" + strconv.Itoa(d.a)
}

func main()  {
	var s string
	var n int
	var d data

	// Menu
	fmt.Println("Modifica la data")
	fmt.Print("Inserisci:\n- g per modigicare il giorno\n- m per modificare il mese\n- a per modificare l'anno\n- q per uscire\n-->")
	for {
		for {
			fmt.Scan(&s)
			if s == "q" {
				return
			}
			if s == "a" || s == "m" || s == "g" {
				break
			}
			fmt.Print("Inserisci un valore valido: ")
		}
		
		// inserimento per modifica
		fmt.Print("Inserisci la modifica al numero: ")
		
		// inserisco la modifica
		loop: for {
			fmt.Scan(&n)
			
			switch s {
			case "a":
				break loop
			case "m":
				if n >= 1 && n <= 12 {
				break loop
				}
			case "g":
				if n >= 1 && n <= 31 {
					break loop
				}
			}
			fmt.Print("Inserisci un valore valido: ")
		}
	
		setData(&d, s, n)
		fmt.Println("nuova data:", stringa(d))
		fmt.Print("nuova modifica: ")
	}
}
