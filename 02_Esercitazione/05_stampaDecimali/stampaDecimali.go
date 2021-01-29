// Esercizio: 5
// Dato un numero, stampare solo la parte decimale 

package main

import (
	"fmt"
)

func main() {
	var num1 float64
	
	fmt.Print("Inserisci un numero con la virgola ti restituirò solo la parte decimale: ")
	fmt.Scan(&num1)
	fmt.Println("La parte frazionaria è", num1 - float64(int(num1)))
}