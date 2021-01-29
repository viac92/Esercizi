// Esercizio.
// Basandosi sull'anno di nascita stabilire se si può votare per la camera(18) o per il senato (25).

package main

import (
	"fmt"
)

func main() {

	var annoNascita, oggi int
	oggi = 2020
	
	fmt.Print("Inserisci il tuo anno di nascita: ")
	fmt.Scan(&annoNascita)
	eta := oggi - annoNascita
	
	if eta >= 18 && eta < 25 {
		fmt.Println("Puoi votare solo per la Camera")
	} else if eta < 18 {
		fmt.Println("Non puoi ancora votare.")
	} else {
		fmt.Println("Puoi votare per la Camera e per il Senato")
	}
}
