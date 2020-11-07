// Un primo di Mersenne è un numero primo della forma 2^x - 1. 
// scrivereun programma che stampi i primi n numeri di Mersenne. 

package main

import (
	"fmt"
	"math"
)

func primiMersenne(n float64) (out int) {
	out = int(math.Pow(2, n)) - 1
	return 
}

func main()  {
	var n int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++{
		fmt.Println(primiMersenne(float64(i)))
	}
	
}