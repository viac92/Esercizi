// inserire due età e valutare chi è più grande

package main

import "fmt"

func main() {
	var piero, luigi int
	fmt.Print("Inserire l'età di Piero: ")
	fmt.Scan(&piero)
	fmt.Print("Inserisci l'età di Luigi: ")
	fmt.Scan(&luigi)

	if piero > luigi {
		fmt.Println("Piero ha", piero, "anni, Luigi ha", luigi, "anni. Piero è più grande di Luigi.")
	} else if piero < luigi {
		fmt.Println("Piero ha", piero, "anni, Luigi ha", luigi, "anni. Luigi è più grande di Piero.")
	} else {
		fmt.Println("Piero ha", piero, "anni, Luigi ha", luigi, "anni. Piero e Luigi hanno la stessa età.")
	}
}