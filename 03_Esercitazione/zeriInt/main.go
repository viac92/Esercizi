// Stabilire se un numero intero contiene qualche zero nelle ultime 3 cifre a destra

package main

import "fmt"

func main() {
	var num1 int
	
	fmt.Print("Inserire un numero intero: ")
	fmt.Scan(&num1)

	if num1 % 10 == 0 && (num1 % 100) - num1 % 10 == 0 && (num1 % 1000) - num1 % 100 == 0 {
		fmt.Println("Il numero", num1, "ha 0 nella prima, seconda e terza posizione.")
	} else if num1 % 10 == 0 && (num1 % 100) - num1 % 10 == 0 {
		fmt.Println("Il numero", num1, "ha 0 nella prima e seconda posizione.")
	} else if num1 % 10 == 0 && (num1 % 1000) - num1 % 100 == 0 {
		fmt.Println("Il numero", num1, "ha 0 nella prima e terza posizione.")
	} else if (num1 % 100) - num1 % 10 == 0 && (num1 % 1000) - num1 % 100 == 0 {
		fmt.Println("Il numero", num1, "ha 0 nella seconda e terza posizione")
	} else if num1 % 10 == 0 && (num1 % 100) - num1 % 10 != 0 && (num1 % 1000) - num1 % 100 != 0 {
		fmt.Println("Il numero", num1, "ha 0 nella prima posizione.")
	} else if num1 % 10 != 0 && (num1 % 100) - num1 % 10 == 0 && (num1 % 1000) - num1 % 100 != 0 {
		fmt.Println("Il numero", num1, "ha 0 nella seconda posizione.")
	} else if num1 % 10 != 0 && (num1 % 100) - num1 % 10 != 0 && (num1 % 1000) - num1 % 100 == 0 {
		fmt.Println("Il numero", num1, "ha 0 nella terza posizione.")
	} else {
		fmt.Println("Il numero", num1, "non ha zeri nelle ultime tre cifre.")
	}
}