// 3 Scrivere un programma che stampi un rettangolo di lunghezza n con la diagonale.

package main

import "fmt"


// funzione che dato n da in uscita una stringa di * quadrata con diagonale
func quadratoDiagonale(n int) string {
	var out string
	var righe int
	//righe
	for righe = 0; righe < n; righe++{
		//colonne
		if righe == 0 || righe == n - 1 {
			for i := 0; i < n; i++ {
				out += "*"
			}
		} else {
			for i := 0; i < n; i++ {
				if i == righe || i == 0 || i == n - 1 {
					out += "*"
				} else {
					out += " "
				}
			}
		}
		out += "\n"
	}
	return out
}


func main()  {
	var n int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)

	fmt.Print(quadratoDiagonale(n))
}