// Esercizio: 1
// Creare un server locale che disegna mediante turtle drawing un'immagine di frattali

package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"math"
	"net/http"
	"github.com/holizz/terrapin"
)

func quadratoHendler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/quadrato.png" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	i := image.NewRGBA(image.Rect(0, 0, 700, 700))
	t := terrapin.NewTerrapin(i, terrapin.Position{300.0, 400.0})
	for i := 0; i < 4; i++ {
		t.Forward(300)
		t.Left(math.Pi / 2)
	}
	png.Encode(w, i)
}
func kochCurve(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/koch.png" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}

	if r.Method != "GET" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}

	i := image.NewRGBA(image.Rect(0, 0, 700, 700))
	t := terrapin.NewTerrapin(i, terrapin.Position{200.0, 500.0})
	for i := 0; i < 4; i++ {
		kochCurveMaker(t, 4.0, 4)
		t.Right(math.Pi - math.Pi / 2)
	}
	png.Encode(w, i)
}

func kochCurveMaker(t *terrapin.Terrapin, lung float64, liv int) {
	if liv == 0 {
		t.Forward(lung)
		return
	}
	kochCurveMaker(t, lung, liv - 1)
	t.Left(math.Pi / 3.0)
	kochCurveMaker(t, lung, liv - 1)
	t.Right(math.Pi - math.Pi / 3.0)
	kochCurveMaker(t, lung, liv - 1)
	t.Left(math.Pi / 3)
	kochCurveMaker(t, lung, liv - 1)
	
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/quadrato.png", quadratoHendler)
	http.HandleFunc("/koch.png", kochCurve)

	fmt.Print("Starting server port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
