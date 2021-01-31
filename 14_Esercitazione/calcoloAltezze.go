// Esercizio: 1
// Scrivere un programma che legge le altezze e calcola e stampa media, varianza e scarto quadtratico.

package main

import	( 
	"fmt"
	"math"
)

// media calcola la media
func media(altezze [100]int, n int) int {
	var somma int
	for i := 0; i < 100; i++ {
		somma += altezze[i]
	}
	return somma / n
}

// varianza calcola la varianza
func varianza(altezze [100]int, media int, n int) float64 {
	var scarto float64
	for i := 0; i < n; i++ {
		scarto += math.Pow(float64(altezze[i]) - float64(media), 2)
	}
	return scarto / float64(n - 1)
}

func main() {
	const K = 100
	var count int
	var altezze [K]int

	fmt.Println("Inserisci un'altezza desiderata (inserisci 0 per terminare): ")
	for i := 0; i < K; i++ {
		fmt.Print(i + 1, "a altezza: ")
		fmt.Scan(&altezze[i])
		if altezze[i] == 0 {
			break
		}
		count++
	}
	fmt.Println("Le altezze inserite sono:")
	for i := 0; i < count; i++{
		fmt.Println(altezze[i])
	}
	fmt.Println("La media è: ")
	fmt.Println(media(altezze, count))
	fmt.Println("La varianza è: ")
	fmt.Println(varianza(altezze, media(altezze, count), count))
}
