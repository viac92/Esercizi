// Scrivere una funzione che date due slice x e y trova quanti elementi di x compaiono in y.

package main

import (
	"fmt"
	"bufio"
	"os"
	"sort"
)

func compareElementXinY(x, y []string) []int {
	var contaElementiOut []int
	var count int
	contaElementiOut = make([]int, len(x))
	for i, elementiX := range x {
		for _, elementiY := range y {
			if elementiX == elementiY {
				count++
			}
			contaElementiOut[i] = count
		}
		count = 0
	}
	return contaElementiOut
}


func sliceElementToCompare(x []string) []string {
	var sliceOut []string

	sort.Slice(x, func(i, j int) bool {
		return x[i] < x[j]
	})

	for i, v := range x {
		if i == 0 || v != x[i -1] {
			sliceOut = append(sliceOut, v)
		}
	}
	return sliceOut
}

func main() {
	var x, y []string
	var ins1, ins2 string


	b1 := bufio.NewScanner(os.Stdin)
	fmt.Print("Inserisci elementi di X: ")
	b1.Scan()
	ins1 = b1.Text()

	for _, char := range ins1 {
		x = append(x, string(char))
	}

	b2 := bufio.NewScanner(os.Stdin)
	fmt.Print("Inserisci elementi di Y: ")
	b2.Scan()
	ins2 = b2.Text()

	for _, char := range ins2 {
		y = append(y, string(char))
	}

	sliceOrdinata := sliceElementToCompare(x)
	sliceElementiDaXinY := compareElementXinY(sliceOrdinata, y)
	
	
	fmt.Println(x)
	fmt.Println(y)

	fmt.Println("Elementi di X in Y:")
	for i := 0; i < len(sliceOrdinata); i++ {
		fmt.Printf("%s --> %d\n", sliceOrdinata[i], sliceElementiDaXinY[i])
	}
}