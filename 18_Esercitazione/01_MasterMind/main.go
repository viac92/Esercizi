//	Implementare il gioco mastermind: data una stringa di parole non conosciuta
//	il giocatore tenta di indovinare, viene fornito il numero di lettere giuste ma al posto sbagliato 
//	e il numero di lettere giuste al posto giusto.

//	Implementare sia dal punto di vista del generatore che dal giocatore.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const lenString int = 5

func creaParole() string {
	var r rune
	var lenStringGenerata string
	rand.Seed(time.Now().UnixNano())
	
	for i := 0; i < lenString; i++ {
		r = rune(rand.Intn(26) + 97)
		lenStringGenerata += string(r)
	}
	return lenStringGenerata
}


func letterePosizioneCorretta(str1, str2 string) (int, int) {
	var subSli1, subSli2 []rune
	r1 := []rune(str1)
	r2 := []rune(str2)

	var count int
	for i := 0; i < lenString; i++ {
		if r1[i] == r2[i] {
			count++
		} else {
			subSli1 = append(subSli1, r1[i])
			subSli2 = append(subSli2, r2[i])
		}
	}
	
	var count2 int
	for _,r := range subSli1 {
		for i := 0; i < len(subSli1); i++ {
			if r == subSli2[i] {
				count2++
				break
			}
		} 
	}		
	return count, count2
}



func main() {
	var strGiocatore string
	var letterePosizione, letterePosSbagliata int

	strDaIndovinare := creaParole()

	//fmt.Println(strDaIndovinare)
	
	fmt.Print("Inserisci una stringa di cinque lettere: ")
	for {	
		fmt.Scan(&strGiocatore)
		if len(strGiocatore) == 5 {
			break
		}
		fmt.Print("Inserisci una stringa valida: ")
	}

	for {
		letterePosizione, letterePosSbagliata = letterePosizioneCorretta(strDaIndovinare, strGiocatore)
		if letterePosizione == lenString {
			break
		}
		fmt.Printf("%d nella posizione corretta, %d lettere giuste posizione sbagliata, riprova!", letterePosizione, letterePosSbagliata)
		for {
			fmt.Print("Insercisci una nuova stringa: ")
			fmt.Scan(&strGiocatore)
			if len(strGiocatore) == lenString {
				break
			}
			fmt.Print("Inserisci una stringa valida: ")
		}
	}
	fmt.Printf("Hai indovinato! La lenString misteriosa %s è uguale al tuo inserimento %s.\n", strDaIndovinare, strGiocatore )
}