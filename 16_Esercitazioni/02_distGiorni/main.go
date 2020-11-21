// Scrivere un programma che riceva come argomenti due date e stampi il numero di giorni passati
//     ./dist 2/1/1956 2/1/1934

package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

// Data contiene le date
type Data struct {
	g, m, a int
}

func stringToData(str string) Data {
	var dataOut Data

	slash := strings.Index(str, "/")
	slash2 := strings.LastIndex(str, "/")
	
	dataOut.g, _ = strconv.Atoi(str[:slash])
	dataOut.m, _ = strconv.Atoi(str[slash+1:slash2])
	dataOut.a, _ = strconv.Atoi(str[slash2+1:])

	return dataOut 
}

func giorniMese(m int, a int) int {
	switch m {
	case 4, 6, 8, 9, 11:
		return 30
	case 2:
		if a % 4 == 0 {
			return 29
		} 
			return 28
		
	default: 
		return 31
	}
}

func distData(data1, data2 Data) int {
	var giorniOut int
	for  i := data1.a; i < data2.a; i++ {
		if i % 4 == 0 {
			giorniOut += 366
		} else {
			giorniOut += 365
		}
	}
	
	if data1.m > data2.m {
		for i, a := data1.m, data1.a; i > data2.m; i-- {
			giorniOut -= giorniMese(i, a)
		}
	} else {
		for i, a := data1.m, data2.a; i < data2.m; i++ {
			giorniOut += giorniMese(i, a)
		}
	}

	giorniOut -= data1.g
	giorniOut += data2.g

	return giorniOut
}

func main() {
	var str1, str2 string
	
	str1 = os.Args[1]
	str2 = os.Args[2]

	data1 := stringToData(str1)
	data2 := stringToData(str2)
	fmt.Println(data1)
	fmt.Println(data2)

	fmt.Println(distData(data1, data2))
}
