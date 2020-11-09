// Scrivere una funzione che data una stringa e un intero ruoti la stringa di quel 
// numero di posizioni. es Tappeto, 3 -> petoTap.

package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)


// dato n scambia le prime n lettere posizionandole alla fine della parola 
func ruotaStringa(n int, s string) (string) {
	var prefix, second string
	for i, char := range s {				
		if i < n {
			prefix += string(char)
		} else {
			second += string(char)
		}
	}
	return second + prefix
}

func main() {
	var s string


	b := bufio.NewScanner(os.Stdin)

	fmt.Print("Inserisci una stringa: ")
	for b.Scan() {
		s = b.Text()
		fmt.Print("Inserisci un numero: ")
		
		for b.Scan() {
			sn := b.Text()
			n, err := strconv.Atoi(sn)

			if err != nil || n > len(s) || n < 0 {
				fmt.Println("Errore --- ripeti inserimento:")
				break
			} 
			
			newS := ruotaStringa(n, s)

			fmt.Println("---->", newS)
			break
		}
		fmt.Print("Inserisci una stringa: ")
	}	
}
