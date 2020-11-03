// estendere fino 999 

package main

import "fmt"

func  main()  {
	var n int
	var s string
	fmt.Print("Inserisci un numero da 0 a 999: ")
	fmt.Scan(&n)

	//casi speciali
	if n == 0 || n >= 10 && n <= 19 {
		
		switch n {
		case 0:
			s = "zero"
		case 10:
			s = "dieci"		
		case 11:
			s = "undici"
		case 12:
			s = "dodici"
		case 13:
			s =	"tredici"
		case 14:
			s = "quattordici"
		case 15:
			s = "quindici"
		case 16:
			s = "sedici"
		case 17:
			s = "diciassette"
		case 18:
			s = "diciotto"
		case 19:		
			s =	"diciannove"	
		} 
	} else {
		// unità
		switch n % 10 {
		case 1:
			s = "uno"
		case 2:
			s = "due"
		case 3:
			s = "tre"
		case 4:
			s = "quattro"
		case 5:
			s = "cinque"
		case 6:
			s = "sei"
		case 7:
			s = "sette"
		case 8:
			s = "otto"
		case 9:
			s = "nove"
		}
	} 
	// decine
	if n % 100 >= 20 && n % 100 <= 29 {
		switch n % 10 {
		case 2, 8:
			s = "vent" + s
		default:
			s =  "venti" + s
		} 
	}
	if n % 100 >= 30 && n % 100 <= 39 {
		switch n % 10 {			
		case 2, 8:
			s = "trent" + s
		default:
			s = "trenta" + s
		} 
	}
	if n % 100 >= 40 && n % 100 <= 49 {
		switch n % 10 {
		case 2, 8:
			s = "quaranta" + s
		default:
			s = "quarant" + s
		} 
	}
	if n % 100 >= 50 && n % 100 <= 59 {
		switch n % 10 {
		case 2, 8:
			s = "cinquant" + s
		default:
			s = "cinquanta" + s
		} 
	}
	if n % 100 >= 60 && n % 100 <= 69 {
		switch n % 10 {
		case 2, 8:
			s = "sesant" + s
		default:
			s = "sesanta" + s
		} 
	}
	if n % 100 >= 70 && n % 100 <= 79 {
		switch n % 10 {
		case 2, 8:
			s = "settant" + s
		default:
			s = "settanta" + s
		} 
	}
	if n % 100 >= 80 && n % 100 <= 89 {
		switch n % 10 {
		case 2, 8:
			s = "ottant" + s
		default:
			s = "ottanta" + s
		} 
	}
	if n % 100 >= 90 && n % 100 <= 99 {
		switch n % 10 {
		case 2, 8:
			s = "novant" + s
		default:
			s = "novanta" + s
		} 
	}
	//centinaia
	if n / 100 == 1 {
		switch (n % 100) / 10 {
		case 8:
			s = "cent" + s 
		default:
			s = "cento" + s
		}
	}
	if n / 100 == 2 {
		switch (n % 100) / 10 {
		case 8:
			s = "duecent" + s 
		default:
			s = "duecento" + s
		}
	}
	if n / 100 == 3 {
		switch (n % 100) / 10 {
		case 8:
			s = "trecent" + s 
		default:
			s = "trecento" + s
		}
	}
	if n / 100 == 4 {
		switch (n % 100) / 10 {
		case 8:
			s = "quattrocent" + s 
		default:
			s = "quattrocento" + s
		}
	}
	if n / 100 == 5 {
		switch (n % 100) / 10 {
		case 8:
			s = "cinquecent" + s 
		default:
			s = "cinquecento" + s
		}
	}
	if n / 100 == 6 {
		switch (n % 100) / 10 {
		case 8:
			s = "seicent" + s 
		default:
			s = "seicento" + s
		}
	}
	if n / 100 == 7 {
		switch (n % 100) / 10 {
		case 8:
			s = "settecent" + s 
		default:
			s = "settecento" + s
		}
	}
	if n / 100 == 8 {
		switch (n % 100) / 10 {
		case 8:
			s = "ottocent" + s 
		default:
			s = "ottocento" + s
		}
	}
	if n / 100 == 9 {
		switch (n % 100) / 10 {
		case 8:
			s = "novecent" + s 
		default:
			s = "novecento" + s
		}
	}
	fmt.Println(s)
}