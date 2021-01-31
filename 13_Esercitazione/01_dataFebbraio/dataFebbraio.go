// Esercizio: 1
// Creare una funzione che restituisca quanti giorni mancano a febbraio usare una struct
package main

import (
		"fmt"
		"strconv"
)

// data è un tipo per salvare date dell'anno
type data struct {
	g, m, a int
}

// isBisestile verifica se un dato anno è bisestile 
func isBisestile(anno int) bool {
	return anno % 4 == 0 && (anno & 100 != 0 || anno % 400 == 0)
}

// ultimoGiornoFebbraio restituisce quanti giorni ha febbraio in un dato anno
func ultimoGiornoDiFebbraio(anno int) data {
	if isBisestile(anno) {
		return data{29, 2, anno}
	}
	return data{28, 2, anno}
}

// stringa restituisce la data in formato string
func stringa(d data) string {
	return strconv.Itoa(d.g) + "/" + strconv.Itoa(d.m) + "/" + strconv.Itoa(d.a)
}

func main()  {
	var n int
	fmt.Print("Inserisci un anno: ")
	fmt.Scan(&n)
	
	s := stringa(ultimoGiornoDiFebbraio(n))
	fmt.Println(s)
}
