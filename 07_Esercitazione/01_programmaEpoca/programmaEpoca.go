// Esercizio: 1
// Sistemare il programma dell'epoca in modo che restituisca valori negativi se prima del 1970.  

package main


import (
	"fmt"
)

func main() {
	var g, m, a int
	
	fmt.Print("Inserisci una data: ")
	fmt.Scan(&g)
	fmt.Scan(&m)
	fmt.Scan(&a)

	fmt.Print(g, "/", m, "/", a)
	fmt.Println()

	gg := 0 
	for anno := 1970; anno < a ; anno++ {
		gg += 365
		if anno % 4 == 0{
			gg++
		}
	}

	for anno := 1970; anno > a; anno-- {
		gg -= 365
		if anno % 4 == 0{
			gg--
		}
	}
	
	for mese := 1; mese < m; mese++ {
		if mese == 11 || mese == 4 || mese == 6 || mese == 9 {
			gg += 30
		} else if mese == 2 {
			gg += 28
			if a % 4 == 0 {
				gg++
			}
		} else {
			gg += 31
		}
	}

	gg += g

	fmt.Print("I giorni trascorsi dal 1/1/1970 al ", g, "/", m, "/", a, " sono passati ", gg, "\n")
}
