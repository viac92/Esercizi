// Leggere una riga e contare quante sono le parole (separati da spazi) [contare gli inizi di parola] 

package main

import (
		"fmt"
		"bufio"
		"os"
)

func main()  {
	var contaParole int
	var flag bool = true
	fmt.Print("Inserisci una frase ti dirò di quante parole è composta: ")
	
	b := bufio.NewScanner(os.Stdin)
	b.Scan()
	s := b.Text()

	for i, x := range s {
		if i == 0 && x == ' ' {
			flag = false
		} 
		if x != ' ' {
			flag = true
		}
		if x == ' ' && flag {
			contaParole++
			flag = false
		}
	}
	if flag == false {
		contaParole--
	}
	fmt.Println("La frase appena inserita contiene:", contaParole + 1, "parole.")
}