/* 		Come programma 1 ma ora il programma chiede di inserire continuamente una nuova linea
		dopo ogni riga stamoare quante "a" conteneva.
	  		. Alla fine stampare la media di tutte le "a" per riga. 
*/

package main

import (
		"fmt"
		"bufio"
		"os"
)

func main() {
	b := bufio.NewScanner(os.Stdin)

	var contaAtotali, contaAparziali, righeTot int
	var media float64
	for {
		fmt.Print("Inserisci una stringa (scrivi exit per uscire): ")
		
		b.Scan()
		s := b.Text()

		if s == "exit" {
			break
		}
		
		righeTot++
		contaAparziali = 0

		for _, x := range s {
			if x == 97 {
				contaAtotali++
				contaAparziali++
			}
		}
		fmt.Println("Ci sono", contaAparziali, " a in questa riga.")
	}
	media = float64(contaAtotali)/float64(righeTot)
	fmt.Println("Il numero totale di \"a\" è:", contaAtotali)
	fmt.Println("La media di \"a\" per riga è di:", media)
	fmt.Println("Numero righe inserite", righeTot)
}