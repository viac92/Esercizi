// Scivere una funzione che date due slice stabilisce se sono uguali. (Stessa lunghezza e stesso contenuto.)

package main

import (
	"fmt"
	"os"
	"bufio"
)

func confrontaSlice(sl1 []string, sl2 []string) bool {
	if len(sl1) != len(sl2) {
		return false
	}

	for i, n := range sl1 {
		if n != sl2[i] {
			return false
		}
	}

	return true
}

func main() {
	var x, y []string
	var inserimento1, inserimento2 string

	fmt.Print("Inserisci una stringa: ")

	b := bufio.NewScanner(os.Stdin)
	
	b.Scan()
	inserimento1 = b.Text()

	for _, char := range inserimento1 {
		x = append(x, string(char))
	}

	b2 := bufio.NewScanner(os.Stdin)
	fmt.Print("Inserisci una seconda stringa: ")
	b2.Scan()
	inserimento2 = b2.Text()

	for _, char := range inserimento2 {
		y = append(y, string(char))
	}

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(confrontaSlice(x, y))
}