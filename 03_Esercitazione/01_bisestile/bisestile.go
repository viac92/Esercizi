// Esercizio: 1
// Stabilire se un anno è bisestile.

package main

import (
	"fmt"
)

func main() {
	var anno int

	fmt.Print("Inserire un anno: ")
	fmt.Scan(&anno)

	if (anno % 100) % 4 == 0 {
		fmt.Println(anno, "è un anno bisestile.")
	} else {
		fmt.Println(anno, "non è un anno bisestile.")
	}
}
