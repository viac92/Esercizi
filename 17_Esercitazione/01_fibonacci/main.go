// Scrivere un programma con una funzione ricorsiva che calcoli gli n numeri di fibonacci.

package main

import "fmt"



func fibonacci(x int) int {
	if x == 0 {
		return 0
	} else if x == 1 || x == 2 {
		return 1
	} else {
		return fibonacci(x-2) + fibonacci(x-1)
	}
}

func main() {
	successiosneFibonacci := [10]int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	
	fmt.Println("N\tFib\tProva")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\t%d\t%d\n", i, successiosneFibonacci[i], fibonacci(i))
	}

	for i := 1; i < 100; i++ {
		fmt.Println(fibonacci(i + 1)/fibonacci(i))
	}
}