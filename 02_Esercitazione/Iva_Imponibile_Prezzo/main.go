// Fare programma che calcola l'iva, imponibile e prezzo.

package main

import "fmt"

func main() {
	// Selettore del programma 
	// Inserire l'operazione desiderata
	var prog int
	fmt.Println("Benvenuto, scegli il programma desiderato:")
	fmt.Print("Calcolo Iva - 1\nCalcolo Imponibile - 2\nCalcolo Prezzo - 3\n")
	fmt.Scan(&prog)

	switch prog {
		case 1: // Iva
			var imp, pFinale float64
			fmt.Print("Inserisci l'imponibile: ")
			fmt.Scan(&imp)
			fmt.Print("Inserisci il prezzo finale: ")
			fmt.Scan(&pFinale)
			fmt.Println("La percentuale di Iva applicata è pari al:", ((100 * pFinale)/imp) - 100, "%")

		case 2: // Imponibile
			var pFinale, iva float64
			fmt.Print("Inserisci il prezzo finale: ")
			fmt.Scan(&pFinale)
			fmt.Print("Inserisci l'iva applicata: ")
			fmt.Scan(&iva)
			fmt.Println("L'imponibile netto è pari a:", (pFinale * 100) / (iva + 100))
		
			case 3:	// Prezzo
			var imp, iva float64
			fmt.Print("Inserisci l'imponibile: ")
			fmt.Scan(&imp)
			fmt.Print("Inserisci l'iva applicata: ")
			fmt.Scan(&iva)
			fmt.Println("Il prezzo finale è pari a:", (imp * (iva * 0.01) + imp) )
			
		default:
			fmt.Println("Operazione non valida.")
	}
}