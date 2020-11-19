// Scrivere un programma che riceva come argomenti due date e stampi il numero di giorni passati
//     ./dist 2/1/1956 2/1/1934

package main

import (
	"fmt"
	"os"
	"strconv"
)

type data struct {
	g, m, a int
}

func stringToData(s string) (data, error) {
	var g, m, a int
	var slash1 int
	var err error
	for i, char := range s {
		if char == '/' && slash1 == 0 {
			slash1 = i
			g,err = strconv.Atoi(s[:i])
		} else if char == '/' {
			m,err = strconv.Atoi(s[slash1 + 1: i])
			a,err = strconv.Atoi(s[i + 1 :])
		}
	}
	return data{g, m, a}, err
}

func dataToString(d data) string {
	return strconv.Itoa(d.g) + "/" + strconv.Itoa(d.m) + "/" + strconv.Itoa(d.a)
}

func distData(d1 data, d2 data) int {
	
}



func main() {
	dataString := os.Args[1]
	dataString2 := os.Args[2]

	fmt.Println(stringToData(dataString))
	fmt.Println(stringToData(dataString2))
}